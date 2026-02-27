package application

import (
	"errors"
	"fmt"
	ws "laundry-hub-api/src/core/websocket"
	machineDomain "laundry-hub-api/src/machine/domain"
	"laundry-hub-api/src/notification/domain"
	notificationEntities "laundry-hub-api/src/notification/domain/entities"
	reservationDomain "laundry-hub-api/src/reservation/domain"
	"laundry-hub-api/src/reservation/domain/entities"
	userDomain "laundry-hub-api/src/user/domain"
	userEntities "laundry-hub-api/src/user/domain/entities"
)

type CreateReservation struct {
	reservationRepo  reservationDomain.IReservationRepository
	machineRepo      machineDomain.IMachineRepository
	notificationRepo domain.INotificationRepository
	userRepo         userDomain.IUserRepository
}

func NewCreateReservation(
	reservationRepo reservationDomain.IReservationRepository,
	machineRepo machineDomain.IMachineRepository,
	notificationRepo domain.INotificationRepository,
	userRepo userDomain.IUserRepository,
) *CreateReservation {
	return &CreateReservation{
		reservationRepo:  reservationRepo,
		machineRepo:      machineRepo,
		notificationRepo: notificationRepo,
		userRepo:         userRepo,
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

	user, err := cr.userRepo.GetByID(userID)
	if err != nil || user == nil {
		user = &userEntities.User{Name: "Usuario"}
	}

	userNotification := &notificationEntities.Notification{
		UserID:        userID,
		ReservationID: &saved.ID,
		Message:       "Tu reservación ha sido creada exitosamente",
		Type:          "RESERVATION_CREATED",
	}
	savedUserNotif, err := cr.notificationRepo.Save(userNotification)
	if err == nil {
		ws.SendNotificationToUser(userID, ws.NotificationPayload{
			ID:            savedUserNotif.ID,
			Message:       savedUserNotif.Message,
			Type:          savedUserNotif.Type,
			ReservationID: savedUserNotif.ReservationID,
		})
	}

	admins, err := cr.userRepo.GetByRole("ADMIN")
	if err == nil {
		for _, admin := range admins {
			adminMessage := fmt.Sprintf("%s %s ha reservado la máquina %d", user.Name, user.PaternalSurname, machineID)
			adminNotification := &notificationEntities.Notification{
				UserID:        admin.ID,
				ReservationID: &saved.ID,
				Message:       adminMessage,
				Type:          "NEW_RESERVATION",
			}
			savedAdminNotif, err := cr.notificationRepo.Save(adminNotification)
			if err == nil {
				ws.SendNotificationToUser(admin.ID, ws.NotificationPayload{
					ID:            savedAdminNotif.ID,
					Message:       savedAdminNotif.Message,
					Type:          savedAdminNotif.Type,
					ReservationID: savedAdminNotif.ReservationID,
				})
			}
		}
	}

	return saved, nil
}
