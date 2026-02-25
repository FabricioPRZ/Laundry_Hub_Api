package domain

import "laundry-hub-api/src/user/domain/entities"

type IUserRepository interface {
	Save(user *entities.User) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	GetByID(id int) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Update(user *entities.User) error
	Delete(id int) error
}
