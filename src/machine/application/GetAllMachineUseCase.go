package application

import (
	"laundry-hub-api/src/machine/domain"
	"laundry-hub-api/src/machine/domain/entities"
)

type GetAllMachines struct {
	machineRepo domain.IMachineRepository
}

func NewGetAllMachines(machineRepo domain.IMachineRepository) *GetAllMachines {
	return &GetAllMachines{machineRepo: machineRepo}
}

func (gm *GetAllMachines) Execute() ([]*entities.Machine, error) {
	return gm.machineRepo.GetAll()
}
