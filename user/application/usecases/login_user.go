package usecases

import (
	"errors"

	"laundry-hub-api/core/security"
	"laundry-hub-api/user/domain"
	"laundry-hub-api/user/domain/entities"
)

type LoginUserUseCase struct {
	userRepo domain.UserRepository
}

func NewLoginUserUseCase(userRepo domain.UserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{
		userRepo: userRepo,
	}
}

func (uc *LoginUserUseCase) Execute(email, password string) (*entities.User, string, error) {
	// Buscar usuario por email
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Verificar contraseña
	if !security.CheckPassword(password, user.Password) {
		return nil, "", errors.New("invalid credentials")
	}

	// Generar token JWT
	token, err := security.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
