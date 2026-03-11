package application

import (
	"errors"
	machineDomain "laundry-hub-api/src/machine/domain"
	"laundry-hub-api/src/maintenance/domain"
	"laundry-hub-api/src/maintenance/domain/entities"
)

type CreateMaintenance struct {
	maintenanceRepo domain.IMaintenanceRepository
	machineRepo     machineDomain.IMachineRepository
}

func NewCreateMaintenance(
	maintenanceRepo domain.IMaintenanceRepository,
	machineRepo machineDomain.IMachineRepository,
) *CreateMaintenance {
	return &CreateMaintenance{maintenanceRepo: maintenanceRepo, machineRepo: machineRepo}
}

func (cm *CreateMaintenance) Execute(machineID int, description string) (*entities.MaintenanceRecord, error) {
	machine, err := cm.machineRepo.GetByID(machineID)
	if err != nil {
		return nil, err
	}
	if machine == nil {
		return nil, errors.New("máquina no encontrada")
	}

	record := &entities.MaintenanceRecord{
		MachineID:   machineID,
		Description: description,
	}

	saved, err := cm.maintenanceRepo.Save(record)
	if err != nil {
		return nil, err
	}

	machine.Status = "MAINTENANCE"
	cm.machineRepo.Update(machine)

	return saved, nil
}