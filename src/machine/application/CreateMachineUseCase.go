package application

import (
	"errors"
	ws "laundry-hub-api/src/core/websocket"
	"laundry-hub-api/src/machine/domain"
	"laundry-hub-api/src/machine/domain/entities"
)

type CreateMachine struct {
	machineRepo domain.IMachineRepository
}

func NewCreateMachine(machineRepo domain.IMachineRepository) *CreateMachine {
	return &CreateMachine{machineRepo: machineRepo}
}

func (cm *CreateMachine) Execute(machine *entities.Machine) (*entities.Machine, error) {
	if machine.Name == "" {
		return nil, errors.New("el nombre es obligatorio")
	}
	if machine.Capacity == "" {
		return nil, errors.New("la capacidad es obligatoria")
	}

	machine.Status = "AVAILABLE"
	saved, err := cm.machineRepo.Save(machine)
	if err != nil {
		return nil, err
	}

	ws.BroadcastNotification(ws.NotificationPayload{
		ID:      0,
		Message: "Una nueva máquina fue agregada",
		Type:    "MACHINE_STATUS_CHANGED",
	})

	return saved, nil
}
