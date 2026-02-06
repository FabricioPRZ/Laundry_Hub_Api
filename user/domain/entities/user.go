package entities

import "time"

// User representa la entidad de usuario en el dominio
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // No se serializa en JSON
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRole define los roles disponibles
const (
	RoleUser  = "USER"
	RoleAdmin = "ADMIN"
)

// IsAdmin verifica si el usuario es administrador
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

// Validate valida los datos del usuario
func (u *User) Validate() error {
	if u.Name == "" {
		return ErrNameRequired
	}
	if u.Email == "" {
		return ErrEmailRequired
	}
	if u.Password == "" {
		return ErrPasswordRequired
	}
	if len(u.Password) < 6 {
		return ErrPasswordTooShort
	}
	return nil
}

// Errores de validación
var (
	ErrNameRequired     = &ValidationError{Message: "Name is required"}
	ErrEmailRequired    = &ValidationError{Message: "Email is required"}
	ErrPasswordRequired = &ValidationError{Message: "Password is required"}
	ErrPasswordTooShort = &ValidationError{Message: "Password must be at least 6 characters"}
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}