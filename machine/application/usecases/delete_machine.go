package usecases

import (
	"errors"

	"laundry-hub-api/machine/domain"
)

type DeleteMachineUseCase struct {
	machineRepo domain.MachineRepository
}

func NewDeleteMachineUseCase(machineRepo domain.MachineRepository) *DeleteMachineUseCase {
	return &DeleteMachineUseCase{
		machineRepo: machineRepo,
	}
}

func (uc *DeleteMachineUseCase) Execute(id string) error {
	if id == "" {
		return errors.New("machine ID is required")
	}

	// Verificar que la máquina existe antes de eliminar
	_, err := uc.machineRepo.FindByID(id)
	if err != nil {
		return err
	}

	return uc.machineRepo.Delete(id)
}
