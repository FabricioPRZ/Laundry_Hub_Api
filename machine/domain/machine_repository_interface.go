package domain

import "laundry-hub-api/machine/domain/entities"

// MachineRepository define la interfaz para el repositorio de máquinas
type MachineRepository interface {
	Create(machine *entities.Machine) error
	FindAll() ([]*entities.Machine, error)
	FindByID(id string) (*entities.Machine, error)
	Update(machine *entities.Machine) error
	Delete(id string) error
}
