package usecases

import (
	"laundry-hub-api/machine/domain"
	"laundry-hub-api/machine/domain/entities"
)

type GetMachinesUseCase struct {
	machineRepo domain.MachineRepository
}

func NewGetMachinesUseCase(machineRepo domain.MachineRepository) *GetMachinesUseCase {
	return &GetMachinesUseCase{
		machineRepo: machineRepo,
	}
}

func (uc *GetMachinesUseCase) Execute() ([]*entities.Machine, error) {
	return uc.machineRepo.FindAll()
}
