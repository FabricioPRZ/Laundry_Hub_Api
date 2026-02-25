package application

import (
	"errors"
	"laundry-hub-api/src/user/domain"
	"laundry-hub-api/src/user/domain/entities"
)

type UpdateUser struct {
	userRepo domain.IUserRepository
}

func NewUpdateUser(userRepo domain.IUserRepository) *UpdateUser {
	return &UpdateUser{userRepo: userRepo}
}

func (uu *UpdateUser) Execute(user *entities.User) error {
	existingUser, err := uu.userRepo.GetByID(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("usuario no encontrado")
	}

	if user.ImageProfile == nil || *user.ImageProfile == "" {
		user.ImageProfile = existingUser.ImageProfile
	}

	return uu.userRepo.Update(user)
}
