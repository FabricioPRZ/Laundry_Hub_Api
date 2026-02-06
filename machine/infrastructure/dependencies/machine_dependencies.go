package dependencies

import (
	"database/sql"

	"laundry-hub-api/machine/application/usecases"
	"laundry-hub-api/machine/infrastructure/adapters"
	"laundry-hub-api/machine/infrastructure/controllers"
)

type MachineDependencies struct {
	GetMachinesController    *controllers.GetMachinesController
	GetMachineByIDController *controllers.GetMachineByIDController
	CreateMachineController  *controllers.CreateMachineController
	UpdateMachineController  *controllers.UpdateMachineController
	DeleteMachineController  *controllers.DeleteMachineController
}

func NewMachineDependencies(db *sql.DB) *MachineDependencies {
	// Adapter
	machineRepo := adapters.NewMachineMySQLRepository(db)

	// Use Cases
	getMachinesUseCase := usecases.NewGetMachinesUseCase(machineRepo)
	getMachineByIDUseCase := usecases.NewGetMachineByIDUseCase(machineRepo)
	createMachineUseCase := usecases.NewCreateMachineUseCase(machineRepo)
	updateMachineUseCase := usecases.NewUpdateMachineUseCase(machineRepo)
	deleteMachineUseCase := usecases.NewDeleteMachineUseCase(machineRepo)

	// Controllers
	getMachinesController := controllers.NewGetMachinesController(getMachinesUseCase)
	getMachineByIDController := controllers.NewGetMachineByIDController(getMachineByIDUseCase)
	createMachineController := controllers.NewCreateMachineController(createMachineUseCase)
	updateMachineController := controllers.NewUpdateMachineController(updateMachineUseCase)
	deleteMachineController := controllers.NewDeleteMachineController(deleteMachineUseCase)

	return &MachineDependencies{
		GetMachinesController:    getMachinesController,
		GetMachineByIDController: getMachineByIDController,
		CreateMachineController:  createMachineController,
		UpdateMachineController:  updateMachineController,
		DeleteMachineController:  deleteMachineController,
	}
}