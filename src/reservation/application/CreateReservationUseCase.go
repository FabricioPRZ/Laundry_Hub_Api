package application

import (
	"errors"
	ws "laundry-hub-api/src/core/websocket"
	machineDomain "laundry-hub-api/src/machine/domain"
	"laundry-hub-api/src/notification/domain"
	notificationEntities "laundry-hub-api/src/notification/domain/entities"
	reservationDomain "laundry-hub-api/src/reservation/domain"
	"laundry-hub-api/src/reservation/domain/entities"
)

type CreateReservation struct {
	reservationRepo  reservationDomain.IReservationRepository
	machineRepo      machineDomain.IMachineRepository
	notificationRepo domain.INotificationRepository
}

func NewCreateReservation(
	reservationRepo reservationDomain.IReservationRepository,
	machineRepo machineDomain.IMachineRepository,
	notificationRepo domain.INotificationRepository,
) *CreateReservation {
	return &CreateReservation{
		reservationRepo:  reservationRepo,
		machineRepo:      machineRepo,
		notificationRepo: notificationRepo,
	}
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

	notification := &notificationEntities.Notification{
		UserID:        userID,
		ReservationID: &saved.ID,
		Message:       "Tu reservación ha sido creada exitosamente",
		Type:          "OCCUPIED",
	}

	savedNotification, err := cr.notificationRepo.Save(notification)
	if err == nil {
		ws.SendNotificationToUser(userID, ws.NotificationPayload{
			ID:            savedNotification.ID,
			Message:       savedNotification.Message,
			Type:          savedNotification.Type,
			ReservationID: savedNotification.ReservationID,
		})
	}

	return saved, nil
}
