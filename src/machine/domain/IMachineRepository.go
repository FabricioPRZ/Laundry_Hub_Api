package domain

import "laundry-hub-api/src/machine/domain/entities"

type IMachineRepository interface {
	Save(machine *entities.Machine) (*entities.Machine, error)
	GetByID(id int) (*entities.Machine, error)
	GetAll() ([]*entities.Machine, error)
	Update(machine *entities.Machine) error
	Delete(id int) error
}
