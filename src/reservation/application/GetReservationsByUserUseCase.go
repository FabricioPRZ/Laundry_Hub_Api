package application

import (
	"laundry-hub-api/src/reservation/domain"
	"laundry-hub-api/src/reservation/domain/entities"
)

type GetReservationsByUser struct {
	reservationRepo domain.IReservationRepository
}

func NewGetReservationsByUser(reservationRepo domain.IReservationRepository) *GetReservationsByUser {
	return &GetReservationsByUser{reservationRepo: reservationRepo}
}

func (gr *GetReservationsByUser) Execute(userID int) ([]*entities.Reservation, error) {
	return gr.reservationRepo.GetByUserID(userID)
}
