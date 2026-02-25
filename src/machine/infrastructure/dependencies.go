package infrastructure

import (
	"laundry-hub-api/src/core"
	"laundry-hub-api/src/machine/application"
	"laundry-hub-api/src/machine/infrastructure/adapters"
	"laundry-hub-api/src/machine/infrastructure/controllers"
)

type DependenciesMachines struct {
	CreateMachineController  *controllers.CreateMachineController
	GetAllMachinesController *controllers.GetAllMachinesController
	GetMachineByIdController *controllers.GetMachineByIdController
	UpdateMachineController  *controllers.UpdateMachineController
	DeleteMachineController  *controllers.DeleteMachineController
}

func InitMachines() *DependenciesMachines {
	conn := core.GetDBPool()
	machineRepo := adapters.NewMySQL(conn.DB)

	createMachine := application.NewCreateMachine(machineRepo)
	getAllMachines := application.NewGetAllMachines(machineRepo)
	getMachineByID := application.NewGetMachineByID(machineRepo)
	updateMachine := application.NewUpdateMachine(machineRepo)
	deleteMachine := application.NewDeleteMachine(machineRepo)

	return &DependenciesMachines{
		CreateMachineController:  controllers.NewCreateMachineController(createMachine),
		GetAllMachinesController: controllers.NewGetAllMachinesController(getAllMachines),
		GetMachineByIdController: controllers.NewGetMachineByIdController(getMachineByID),
		UpdateMachineController:  controllers.NewUpdateMachineController(updateMachine),
		DeleteMachineController:  controllers.NewDeleteMachineController(deleteMachine),
	}
}
