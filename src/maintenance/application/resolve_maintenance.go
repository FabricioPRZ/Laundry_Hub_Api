package application

import (
	"errors"
	machineDomain "laundry-hub-api/src/machine/domain"
	"laundry-hub-api/src/maintenance/domain"
)

type ResolveMaintenance struct {
	maintenanceRepo domain.IMaintenanceRepository
	machineRepo     machineDomain.IMachineRepository
}

func NewResolveMaintenance(
	maintenanceRepo domain.IMaintenanceRepository,
	machineRepo machineDomain.IMachineRepository,
) *ResolveMaintenance {
	return &ResolveMaintenance{maintenanceRepo: maintenanceRepo, machineRepo: machineRepo}
}

func (rm *ResolveMaintenance) Execute(id int) error {
	record, err := rm.maintenanceRepo.GetByID(id)
	if err != nil {
		return err
	}
	if record == nil {
		return errors.New("registro no encontrado")
	}
	if record.IsResolved {
		return errors.New("el registro ya está resuelto")
	}

	if err := rm.maintenanceRepo.Resolve(id); err != nil {
		return err
	}

	machine, err := rm.machineRepo.GetByID(record.MachineID)
	if err != nil {
		return err
	}
	if machine != nil {
		machine.Status = "AVAILABLE"
		rm.machineRepo.Update(machine)
	}

	return nil
}