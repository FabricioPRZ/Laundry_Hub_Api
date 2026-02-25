package application

import (
	"errors"
	"laundry-hub-api/src/core/cloudinary"
	"laundry-hub-api/src/user/domain"
)

type DeleteUser struct {
	userRepo domain.IUserRepository
}

func NewDeleteUser(userRepo domain.IUserRepository) *DeleteUser {
	return &DeleteUser{userRepo: userRepo}
}

func (du *DeleteUser) Execute(id int) error {
	existingUser, err := du.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("usuario no encontrado")
	}

	if existingUser.ImageProfile != nil && *existingUser.ImageProfile != "" {
		if err := cloudinary.DeleteImage(*existingUser.ImageProfile); err != nil {
			return errors.New("error al eliminar imagen de Cloudinary")
		}
	}

	return du.userRepo.Delete(id)
}
