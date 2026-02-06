package usecases

import (
	"laundry-hub-api/user/domain"
	"laundry-hub-api/user/domain/entities"
)

type GetUserByEmailUseCase struct {
	userRepo domain.UserRepository
}

func NewGetUserByEmailUseCase(userRepo domain.UserRepository) *GetUserByEmailUseCase {
	return &GetUserByEmailUseCase{
		userRepo: userRepo,
	}
}

func (uc *GetUserByEmailUseCase) Execute(email string) (*entities.User, error) {
	return uc.userRepo.FindByEmail(email)
}
