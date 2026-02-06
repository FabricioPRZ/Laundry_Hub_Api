package usecases

import (
	"errors"

	"laundry-hub-api/machine/domain"
	"laundry-hub-api/machine/domain/entities"
)

type GetMachineByIDUseCase struct {
	machineRepo domain.MachineRepository
}

func NewGetMachineByIDUseCase(machineRepo domain.MachineRepository) *GetMachineByIDUseCase {
	return &GetMachineByIDUseCase{
		machineRepo: machineRepo,
	}
}

func (uc *GetMachineByIDUseCase) Execute(id string) (*entities.Machine, error) {
	if id == "" {
		return nil, errors.New("machine ID is required")
	}

	return uc.machineRepo.FindByID(id)
}
