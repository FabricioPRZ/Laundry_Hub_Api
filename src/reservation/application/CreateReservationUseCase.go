package application

import (
	"errors"
	machineDomain "laundry-hub-api/src/machine/domain"
	"laundry-hub-api/src/reservation/domain"
	"laundry-hub-api/src/reservation/domain/entities"
)

type CreateReservation struct {
	reservationRepo domain.IReservationRepository
	machineRepo     machineDomain.IMachineRepository
}

func NewCreateReservation(reservationRepo domain.IReservationRepository, machineRepo machineDomain.IMachineRepository) *CreateReservation {
	return &CreateReservation{reservationRepo: reservationRepo, machineRepo: machineRepo}
}

func (cr *CreateReservation) Execute(userID, machineID int) (*entities.Reservation, error) {
	machine, err := cr.machineRepo.GetByID(machineID)
	if err != nil {
		return nil, err
	}
	if machine == nil {
		return nil, errors.New("máquina no encontrada")
	}
	if machine.Status != "AVAILABLE" {
		return nil, errors.New("la máquina no está disponible")
	}

	reservation := &entities.Reservation{
		UserID:    userID,
		MachineID: machineID,
		Status:    "ACTIVE",
	}

	saved, err := cr.reservationRepo.Save(reservation)
	if err != nil {
		return nil, err
	}

	machine.Status = "OCCUPIED"
	if err := cr.machineRepo.Update(machine); err != nil {
		return nil, err
	}

	return saved, nil
}
