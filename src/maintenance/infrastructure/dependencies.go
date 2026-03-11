package infrastructure

import (
	"laundry-hub-api/src/core"
	machineAdapters "laundry-hub-api/src/machine/infrastructure/adapters"
	"laundry-hub-api/src/maintenance/application"
	"laundry-hub-api/src/maintenance/infrastructure/adapters"
	"laundry-hub-api/src/maintenance/infrastructure/controllers"
)

type DependenciesMaintenance struct {
	CreateMaintenanceController  *controllers.CreateMaintenanceController
	GetAllMaintenanceController  *controllers.GetAllMaintenanceController
	ResolveMaintenanceController *controllers.ResolveMaintenanceController
	DeleteMaintenanceController  *controllers.DeleteMaintenanceController
}

func InitMaintenance() *DependenciesMaintenance {
	conn            := core.GetDBPool()
	maintenanceRepo := adapters.NewMySQL(conn.DB)
	machineRepo     := machineAdapters.NewMySQL(conn.DB)

	createMaintenance  := application.NewCreateMaintenance(maintenanceRepo, machineRepo)
	getAllMaintenance   := application.NewGetAllMaintenance(maintenanceRepo)
	resolveMaintenance := application.NewResolveMaintenance(maintenanceRepo, machineRepo)
	deleteMaintenance  := application.NewDeleteMaintenance(maintenanceRepo)

	return &DependenciesMaintenance{
		CreateMaintenanceController:  controllers.NewCreateMaintenanceController(createMaintenance),
		GetAllMaintenanceController:  controllers.NewGetAllMaintenanceController(getAllMaintenance),
		ResolveMaintenanceController: controllers.NewResolveMaintenanceController(resolveMaintenance),
		DeleteMaintenanceController:  controllers.NewDeleteMaintenanceController(deleteMaintenance),
	}
}