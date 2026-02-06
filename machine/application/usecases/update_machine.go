package usecases

import (
	"errors"
	"time"

	"laundry-hub-api/machine/domain"
	"laundry-hub-api/machine/domain/entities"
)

type UpdateMachineUseCase struct {
	machineRepo domain.MachineRepository
}

func NewUpdateMachineUseCase(machineRepo domain.MachineRepository) *UpdateMachineUseCase {
	return &UpdateMachineUseCase{
		machineRepo: machineRepo,
	}
}

func (uc *UpdateMachineUseCase) Execute(id, name, capacity, location, status string) (*entities.Machine, error) {
	if id == "" {
		return nil, errors.New("machine ID is required")
	}

	// Verificar que la máquina existe
	existingMachine, err := uc.machineRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Actualizar campos
	if name != "" {
		existingMachine.Name = name
	}
	if capacity != "" {
		existingMachine.Capacity = capacity
	}
	if location != "" {
		existingMachine.Location = location
	}
	if status != "" {
		existingMachine.Status = status
	}
	existingMachine.UpdatedAt = time.Now()

	// Validar máquina
	if err := existingMachine.Validate(); err != nil {
		return nil, err
	}

	// Actualizar en base de datos
	if err := uc.machineRepo.Update(existingMachine); err != nil {
		return nil, err
	}

	return existingMachine, nil
}
