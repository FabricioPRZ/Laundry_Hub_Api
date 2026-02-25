package application

import (
	"laundry-hub-api/src/machine/domain"
	"laundry-hub-api/src/machine/domain/entities"
)

type GetMachineByID struct {
	machineRepo domain.IMachineRepository
}

func NewGetMachineByID(machineRepo domain.IMachineRepository) *GetMachineByID {
	return &GetMachineByID{machineRepo: machineRepo}
}

func (gm *GetMachineByID) Execute(id int) (*entities.Machine, error) {
	return gm.machineRepo.GetByID(id)
}
