package application

import (
	"errors"
	ws "laundry-hub-api/src/core/websocket"
	machineDomain "laundry-hub-api/src/machine/domain"
	"laundry-hub-api/src/maintenance/domain"
	"laundry-hub-api/src/maintenance/domain/entities"
	notificationDomain "laundry-hub-api/src/notification/domain"
	notificationEntities "laundry-hub-api/src/notification/domain/entities"
)

type CreateMaintenance struct {
	maintenanceRepo  domain.IMaintenanceRepository
	machineRepo      machineDomain.IMachineRepository
	notificationRepo notificationDomain.INotificationRepository
}

func NewCreateMaintenance(
	maintenanceRepo  domain.IMaintenanceRepository,
	machineRepo      machineDomain.IMachineRepository,
	notificationRepo notificationDomain.INotificationRepository,
) *CreateMaintenance {
	return &CreateMaintenance{
		maintenanceRepo:  maintenanceRepo,
		machineRepo:      machineRepo,
		notificationRepo: notificationRepo,
	}
}

func (cm *CreateMaintenance) Execute(userID, machineID int, description string) (*entities.MaintenanceRecord, error) {
	machine, err := cm.machineRepo.GetByID(machineID)
	if err != nil {
		return nil, err
	}
	if machine == nil {
		return nil, errors.New("máquina no encontrada")
	}

	record := &entities.MaintenanceRecord{
		MachineID:   machineID,
		Description: description,
	}

	saved, err := cm.maintenanceRepo.Save(record)
	if err != nil {
		return nil, err
	}

	machine.Status = "MAINTENANCE"
	cm.machineRepo.Update(machine)

	notification := &notificationEntities.Notification{
		UserID:  userID,
		Message: "Máquina \"" + machine.Name + "\" puesta en mantenimiento",
		Type:    "MAINTENANCE",
	}
	savedNotification, notifErr := cm.notificationRepo.Save(notification)

	payload := ws.NotificationPayload{
		Message: "Máquina \"" + machine.Name + "\" puesta en mantenimiento",
		Type:    "MAINTENANCE",
	}
	if notifErr == nil {
		payload.ID = savedNotification.ID
	}
	ws.BroadcastNotification(payload)

	return saved, nil
}