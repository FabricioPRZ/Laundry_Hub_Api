package application

import (
	"laundry-hub-api/src/user/domain"
	"laundry-hub-api/src/user/domain/entities"
)

type GetAllUsers struct {
	userRepo domain.IUserRepository
}

func NewGetAllUsers(userRepo domain.IUserRepository) *GetAllUsers {
	return &GetAllUsers{userRepo: userRepo}
}

func (gu *GetAllUsers) Execute() ([]*entities.User, error) {
	return gu.userRepo.GetAll()
}
