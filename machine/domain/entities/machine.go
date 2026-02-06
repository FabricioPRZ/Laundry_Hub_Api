package entities

import "time"

// Machine representa la entidad de máquina en el dominio
type Machine struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Capacity  string    `json:"capacity"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MachineStatus define los estados disponibles
const (
	StatusAvailable   = "AVAILABLE"
	StatusOccupied    = "OCCUPIED"
	StatusMaintenance = "MAINTENANCE"
)

// IsAvailable verifica si la máquina está disponible
func (m *Machine) IsAvailable() bool {
	return m.Status == StatusAvailable
}

// Validate valida los datos de la máquina
func (m *Machine) Validate() error {
	if m.Name == "" {
		return ErrNameRequired
	}
	if m.Capacity == "" {
		return ErrCapacityRequired
	}
	if m.Location == "" {
		return ErrLocationRequired
	}
	if !isValidStatus(m.Status) {
		return ErrInvalidStatus
	}
	return nil
}

func isValidStatus(status string) bool {
	return status == StatusAvailable || status == StatusOccupied || status == StatusMaintenance
}

// Errores de validación
var (
	ErrNameRequired     = &ValidationError{Message: "Name is required"}
	ErrCapacityRequired = &ValidationError{Message: "Capacity is required"}
	ErrLocationRequired = &ValidationError{Message: "Location is required"}
	ErrInvalidStatus    = &ValidationError{Message: "Invalid status"}
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}