package infrastructure

import (
	"laundry-hub-api/src/core"
	machineAdapters "laundry-hub-api/src/machine/infrastructure/adapters"
	"laundry-hub-api/src/maintenance/application"
	maintenanceAdapters "laundry-hub-api/src/maintenance/infrastructure/adapters"
	"laundry-hub-api/src/maintenance/infrastructure/controllers"
	notificationAdapters "laundry-hub-api/src/notification/infrastructure/adapters"
)

type DependenciesMaintenance struct {
	CreateMaintenanceController  *controllers.CreateMaintenanceController
	GetAllMaintenanceController  *controllers.GetAllMaintenanceController
	ResolveMaintenanceController *controllers.ResolveMaintenanceController
	DeleteMaintenanceController  *controllers.DeleteMaintenanceController
}

func InitMaintenance() *DependenciesMaintenance {
	conn             := core.GetDBPool()
	maintenanceRepo  := maintenanceAdapters.NewMySQL(conn.DB)
	machineRepo      := machineAdapters.NewMySQL(conn.DB)
	notificationRepo := notificationAdapters.NewMySQL(conn.DB)

	createMaintenance  := application.NewCreateMaintenance(maintenanceRepo, machineRepo, notificationRepo)
	getAllMaintenance   := application.NewGetAllMaintenance(maintenanceRepo)
	resolveMaintenance := application.NewResolveMaintenance(maintenanceRepo, machineRepo, notificationRepo)
	deleteMaintenance  := application.NewDeleteMaintenance(maintenanceRepo)

	return &DependenciesMaintenance{
		CreateMaintenanceController:  controllers.NewCreateMaintenanceController(createMaintenance),
		GetAllMaintenanceController:  controllers.NewGetAllMaintenanceController(getAllMaintenance),
		ResolveMaintenanceController: controllers.NewResolveMaintenanceController(resolveMaintenance),
		DeleteMaintenanceController:  controllers.NewDeleteMaintenanceController(deleteMaintenance),
	}
}