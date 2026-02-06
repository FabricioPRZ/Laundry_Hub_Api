package usecases

import (
	"errors"
	"time"

	"laundry-hub-api/core/security"
	"laundry-hub-api/user/domain"
	"laundry-hub-api/user/domain/entities"

	"github.com/google/uuid"
)

type RegisterUserUseCase struct {
	userRepo domain.UserRepository
}

func NewRegisterUserUseCase(userRepo domain.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userRepo: userRepo,
	}
}

func (uc *RegisterUserUseCase) Execute(name, email, password string) (*entities.User, error) {
	// Validar que el email no esté ya registrado
	existingUser, _ := uc.userRepo.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hashear contraseña
	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Crear usuario
	user := &entities.User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		Role:      entities.RoleUser,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Validar usuario
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// Guardar en base de datos
	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
