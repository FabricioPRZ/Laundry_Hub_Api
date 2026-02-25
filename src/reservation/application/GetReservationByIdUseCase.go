package application

import (
	"laundry-hub-api/src/reservation/domain"
	"laundry-hub-api/src/reservation/domain/entities"
)

type GetReservationByID struct {
	reservationRepo domain.IReservationRepository
}

func NewGetReservationByID(reservationRepo domain.IReservationRepository) *GetReservationByID {
	return &GetReservationByID{reservationRepo: reservationRepo}
}

func (gr *GetReservationByID) Execute(id int) (*entities.Reservation, error) {
	return gr.reservationRepo.GetByID(id)
}
