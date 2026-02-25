package domain

import (
	"laundry-hub-api/src/reservation/domain/entities"
	"time"
)

type IReservationRepository interface {
	Save(reservation *entities.Reservation) (*entities.Reservation, error)
	GetByID(id int) (*entities.Reservation, error)
	GetByUserID(userID int) ([]*entities.Reservation, error)
	UpdateStatus(id int, status string, endedAt *time.Time) error
}
