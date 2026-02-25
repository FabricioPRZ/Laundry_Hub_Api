package application

import (
	"errors"
	"laundry-hub-api/src/user/domain"
	"laundry-hub-api/src/user/domain/entities"
)

type CreateUser struct {
	userRepo domain.IUserRepository
}

func NewCreateUser(userRepo domain.IUserRepository) *CreateUser {
	return &CreateUser{userRepo: userRepo}
}

func (cu *CreateUser) Execute(user *entities.User) (*entities.User, error) {
	if user.Name == "" {
		return nil, errors.New("el nombre es obligatorio")
	}
	if user.Email == "" {
		return nil, errors.New("el email es obligatorio")
	}
	if user.Password == nil || *user.Password == "" {
		return nil, errors.New("la contraseña es obligatoria")
	}

	existingUser, _ := cu.userRepo.GetByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("el email ya está registrado")
	}

	return cu.userRepo.Save(user)
}
