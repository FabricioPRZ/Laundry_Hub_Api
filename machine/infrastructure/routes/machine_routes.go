package routes

import (
	"github.com/gin-gonic/gin"
	"laundry-hub-api/core/security"
	"laundry-hub-api/machine/infrastructure/dependencies"
)

func RegisterMachineRoutes(api *gin.RouterGroup, deps *dependencies.MachineDependencies) {
	machines := api.Group("/machines")
	{
		// Rutas públicas (pueden ver las máquinas)
		machines.GET("", deps.GetMachinesController.Handle)
		machines.GET("/:id", deps.GetMachineByIDController.Handle)

		// Rutas protegidas - Solo Admin
		protected := machines.Group("")
		protected.Use(security.AuthMiddleware(), security.AdminMiddleware())
		{
			protected.POST("", deps.CreateMachineController.Handle)
			protected.PUT("/:id", deps.UpdateMachineController.Handle)
			protected.DELETE("/:id", deps.DeleteMachineController.Handle)
		}
	}
}