package application

import (
	"errors"
	ws "laundry-hub-api/src/core/websocket"
	"laundry-hub-api/src/machine/domain"
	"laundry-hub-api/src/machine/domain/entities"
)

type UpdateMachine struct {
	machineRepo domain.IMachineRepository
}

func NewUpdateMachine(machineRepo domain.IMachineRepository) *UpdateMachine {
	return &UpdateMachine{machineRepo: machineRepo}
}

func (um *UpdateMachine) Execute(machine *entities.Machine) error {
	existing, err := um.machineRepo.GetByID(machine.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("máquina no encontrada")
	}

	if err := um.machineRepo.Update(machine); err != nil {
		return err
	}

	ws.BroadcastNotification(ws.NotificationPayload{
		ID:      0,
		Message: "Una máquina fue actualizada",
		Type:    "MACHINE_STATUS_CHANGED",
	})

	return nil
}
