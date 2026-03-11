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

func (rm *ResolveMaintenance) Execute(id int) error {
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
			UserID:  0,
			Message: "Máquina \"" + machine.Name + "\" disponible nuevamente",
			Type:    "RELEASED",
		}
		savedNotification, err := rm.notificationRepo.Save(notification)
		if err == nil {
			ws.BroadcastNotification(ws.NotificationPayload{
				ID:      savedNotification.ID,
				Message: savedNotification.Message,
				Type:    savedNotification.Type,
			})
		}
	}

	return nil
}