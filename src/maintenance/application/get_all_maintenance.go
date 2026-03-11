package application

import (
	"laundry-hub-api/src/maintenance/domain"
	"laundry-hub-api/src/maintenance/domain/entities"
)

type GetAllMaintenance struct {
	maintenanceRepo domain.IMaintenanceRepository
}

func NewGetAllMaintenance(repo domain.IMaintenanceRepository) *GetAllMaintenance {
	return &GetAllMaintenance{maintenanceRepo: repo}
}

func (g *GetAllMaintenance) Execute() ([]*entities.MaintenanceRecord, error) {
	return g.maintenanceRepo.GetAll()
}