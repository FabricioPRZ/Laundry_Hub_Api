package application

import (
	"errors"
	machineDomain "laundry-hub-api/src/machine/domain"
	"laundry-hub-api/src/reservation/domain"
	"time"
)

type CancelReservation struct {
	reservationRepo domain.IReservationRepository
	machineRepo     machineDomain.IMachineRepository
}

func NewCancelReservation(reservationRepo domain.IReservationRepository, machineRepo machineDomain.IMachineRepository) *CancelReservation {
	return &CancelReservation{reservationRepo: reservationRepo, machineRepo: machineRepo}
}

func (cr *CancelReservation) Execute(reservationID, userID int) error {
	reservation, err := cr.reservationRepo.GetByID(reservationID)
	if err != nil {
		return err
	}
	if reservation == nil {
		return errors.New("reservación no encontrada")
	}
	if reservation.UserID != userID {
		return errors.New("no tienes permiso para cancelar esta reservación")
	}
	if reservation.Status != "ACTIVE" {
		return errors.New("solo se pueden cancelar reservaciones activas")
	}

	now := time.Now()
	if err := cr.reservationRepo.UpdateStatus(reservationID, "CANCELLED", &now); err != nil {
		return err
	}

	machine, err := cr.machineRepo.GetByID(reservation.MachineID)
	if err != nil {
		return err
	}
	if machine != nil {
		machine.Status = "AVAILABLE"
		cr.machineRepo.Update(machine)
	}

	return nil
}
