package routes

import (
	"laundry-hub-api/src/core/security"
	"laundry-hub-api/src/machine/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureMachineRoutes(
	router *gin.Engine,
	createCtrl *controllers.CreateMachineController,
	getAllCtrl *controllers.GetAllMachinesController,
	getByIDCtrl *controllers.GetMachineByIdController,
	updateCtrl *controllers.UpdateMachineController,
	deleteCtrl *controllers.DeleteMachineController,
) {
	machineGroup := router.Group("/machines")
	machineGroup.Use(security.JWTMiddleware())
	{
		machineGroup.POST("", createCtrl.Execute)
		machineGroup.GET("", getAllCtrl.Execute)
		machineGroup.GET("/:id", getByIDCtrl.Execute)
		machineGroup.PUT("/:id", updateCtrl.Execute)
		machineGroup.DELETE("/:id", deleteCtrl.Execute)
	}
}
