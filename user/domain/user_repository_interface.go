package domain

import "laundry-hub-api/user/domain/entities"

// UserRepository define la interfaz para el repositorio de usuarios
type UserRepository interface {
	Create(user *entities.User) error
	FindByID(id string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id string) error
}
