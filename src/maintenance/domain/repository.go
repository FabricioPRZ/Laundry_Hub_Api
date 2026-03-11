package domain

import "laundry-hub-api/src/maintenance/domain/entities"

type IMaintenanceRepository interface {
	Save(record *entities.MaintenanceRecord) (*entities.MaintenanceRecord, error)
	GetAll() ([]*entities.MaintenanceRecord, error)
	GetByID(id int) (*entities.MaintenanceRecord, error)
	Resolve(id int) error
	Delete(id int) error
}