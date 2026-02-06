package routes

import (
    "github.com/gin-gonic/gin"
    "laundry-hub-api/machine/infrastructure/dependencies"
)

func RegisterMachineRoutes(api *gin.RouterGroup, deps *dependencies.MachineDependencies) {
    machines := api.Group("/machines")
    {        
        // Obtener máquinas (públicas)
        machines.GET("", deps.GetMachinesController.Handle)
        machines.GET("/:id", deps.GetMachineByIDController.Handle)

        machines.POST("", deps.CreateMachineController.Handle)
        machines.PUT("/:id", deps.UpdateMachineController.Handle)
        machines.DELETE("/:id", deps.DeleteMachineController.Handle)
    }
}