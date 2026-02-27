package application

import (
	"errors"
	ws "laundry-hub-api/src/core/websocket"
	"laundry-hub-api/src/machine/domain"
)

type DeleteMachine struct {
	machineRepo domain.IMachineRepository
}

func NewDeleteMachine(machineRepo domain.IMachineRepository) *DeleteMachine {
	return &DeleteMachine{machineRepo: machineRepo}
}

func (dm *DeleteMachine) Execute(id int) error {
	existing, err := dm.machineRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("máquina no encontrada")
	}

	if err := dm.machineRepo.Delete(id); err != nil {
		return err
	}

	ws.BroadcastNotification(ws.NotificationPayload{
		ID:      0,
		Message: "Una máquina fue eliminada",
		Type:    "MACHINE_STATUS_CHANGED",
	})

	return nil
}
