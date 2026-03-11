package application

import (
	"errors"
	ws "laundry-hub-api/src/core/websocket"
	machineDomain "laundry-hub-api/src/machine/domain"
	"laundry-hub-api/src/maintenance/domain"
	notificationDomain "laundry-hub-api/src/notification/domain"
	notificationEntities "laundry-hub-api/src/notification/domain/entities"
)

type ResolveMaintenance struct {
	maintenanceRepo  domain.IMaintenanceRepository
	machineRepo      machineDomain.IMachineRepository
	notificationRepo notificationDomain.INotificationRepository
}

func NewResolveMaintenance(
	maintenanceRepo  domain.IMaintenanceRepository,
	machineRepo      machineDomain.IMachineRepository,
	notificationRepo notificationDomain.INotificationRepository,
) *ResolveMaintenance {
	return &ResolveMaintenance{
		maintenanceRepo:  maintenanceRepo,
		machineRepo:      machineRepo,
		notificationRepo: notificationRepo,
	}
}

func (rm *ResolveMaintenance) Execute(id, userID int) error {
	record, err := rm.maintenanceRepo.GetByID(id)
	if err != nil {
		return err
	}
	if record == nil {
		return errors.New("registro no encontrado")
	}
	if record.IsResolved {
		return errors.New("el registro ya está resuelto")
	}

	if err := rm.maintenanceRepo.Resolve(id); err != nil {
		return err
	}

	machine, err := rm.machineRepo.GetByID(record.MachineID)
	if err != nil {
		return err
	}
	if machine != nil {
		machine.Status = "AVAILABLE"
		rm.machineRepo.Update(machine)

		notification := &notificationEntities.Notification{
			UserID:  userID,
			Message: "Máquina \"" + machine.Name + "\" disponible nuevamente",
			Type:    "RELEASED",
		}
		savedNotification, notifErr := rm.notificationRepo.Save(notification)

		payload := ws.NotificationPayload{
			Message: "Máquina \"" + machine.Name + "\" disponible nuevamente",
			Type:    "RELEASED",
		}
		if notifErr == nil {
			payload.ID = savedNotification.ID
		}
		ws.BroadcastNotification(payload)
	}

	return nil
}