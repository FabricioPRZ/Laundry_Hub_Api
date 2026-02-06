package usecases

import (
	"laundry-hub-api/user/domain"
	"laundry-hub-api/user/domain/entities"
)

type GetUserByIDUseCase struct {
	userRepo domain.UserRepository
}

func NewGetUserByIDUseCase(userRepo domain.UserRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{
		userRepo: userRepo,
	}
}

func (uc *GetUserByIDUseCase) Execute(id string) (*entities.User, error) {
	return uc.userRepo.FindByID(id)
}
