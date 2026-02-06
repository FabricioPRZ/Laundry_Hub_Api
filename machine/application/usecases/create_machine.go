package usecases

import (
	"time"

	"laundry-hub-api/machine/domain"
	"laundry-hub-api/machine/domain/entities"

	"github.com/google/uuid"
)

type CreateMachineUseCase struct {
	machineRepo domain.MachineRepository
}

func NewCreateMachineUseCase(machineRepo domain.MachineRepository) *CreateMachineUseCase {
	return &CreateMachineUseCase{
		machineRepo: machineRepo,
	}
}

func (uc *CreateMachineUseCase) Execute(name, capacity, location, status string) (*entities.Machine, error) {
	// Crear máquina
	machine := &entities.Machine{
		ID:        uuid.New().String(),
		Name:      name,
		Status:    status,
		Capacity:  capacity,
		Location:  location,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Validar máquina
	if err := machine.Validate(); err != nil {
		return nil, err
	}

	// Guardar en base de datos
	if err := uc.machineRepo.Create(machine); err != nil {
		return nil, err
	}

	return machine, nil
}
