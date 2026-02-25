package application

import (
	"errors"
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

	return dm.machineRepo.Delete(id)
}
