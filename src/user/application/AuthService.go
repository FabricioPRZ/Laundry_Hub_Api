package application

import (
	"errors"
	"fmt"
	"laundry-hub-api/src/core/security"
	"laundry-hub-api/src/user/domain"
	"laundry-hub-api/src/user/domain/entities"
	"strings"
)

type AuthService struct {
	userRepo domain.IUserRepository
}

func NewAuthService(userRepo domain.IUserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (as *AuthService) Login(email, password string) (*entities.User, error) {
	email = strings.TrimSpace(email)
	fmt.Println("Buscando usuario con correo:", email)

	user, err := as.userRepo.GetByEmail(email)
	if err != nil {
		fmt.Println("Error al obtener usuario:", err)
		return nil, errors.New("error al buscar usuario")
	}

	if user == nil {
		fmt.Println("Usuario no encontrado")
		return nil, errors.New("credenciales inválidas")
	}

	if user.Password == nil {
		return nil, errors.New("usuario registrado con OAuth - usa login social")
	}

	if !security.CheckPassword(*user.Password, password) {
		fmt.Println("Contraseña incorrecta")
		return nil, errors.New("credenciales inválidas")
	}

	return user, nil
}

func (as *AuthService) Register(user *entities.User) (*entities.User, error) {
	existingUser, _ := as.userRepo.GetByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("el email ya está registrado")
	}

	if user.Password == nil || *user.Password == "" {
		return nil, errors.New("la contraseña es requerida")
	}

	hashedPassword, err := security.HashPassword(*user.Password)
	if err != nil {
		return nil, errors.New("error al procesar la contraseña")
	}

	user.Password = &hashedPassword
	user.OAuthProvider = "LOCAL"
	user.Role = "USER"

	savedUser, err := as.userRepo.Save(user)
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}

func (as *AuthService) GetUserByID(userID int) (*entities.User, error) {
	user, err := as.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("usuario no encontrado")
	}
	return user, nil
}
