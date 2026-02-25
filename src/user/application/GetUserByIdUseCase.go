package application

import (
	"laundry-hub-api/src/user/domain"
	"laundry-hub-api/src/user/domain/entities"
)

type GetUserById struct {
	userRepo domain.IUserRepository
}

func NewGetUserById(userRepo domain.IUserRepository) *GetUserById {
	return &GetUserById{userRepo: userRepo}
}

func (gu *GetUserById) Execute(id int) (*entities.User, error) {
	return gu.userRepo.GetByID(id)
}
