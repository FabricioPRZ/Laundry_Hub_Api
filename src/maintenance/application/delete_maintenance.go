package application

import (
	"errors"
	"laundry-hub-api/src/maintenance/domain"
)

type DeleteMaintenance struct {
	maintenanceRepo domain.IMaintenanceRepository
}

func NewDeleteMaintenance(repo domain.IMaintenanceRepository) *DeleteMaintenance {
	return &DeleteMaintenance{maintenanceRepo: repo}
}

func (dm *DeleteMaintenance) Execute(id int) error {
	record, err := dm.maintenanceRepo.GetByID(id)
	if err != nil {
		return err
	}
	if record == nil {
		return errors.New("registro no encontrado")
	}
	return dm.maintenanceRepo.Delete(id)
}