package routes

import (
	"laundry-hub-api/src/core/security"
	"laundry-hub-api/src/maintenance/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureMaintenanceRoutes(
	router       *gin.Engine,
	createCtrl   *controllers.CreateMaintenanceController,
	getAllCtrl    *controllers.GetAllMaintenanceController,
	resolveCtrl  *controllers.ResolveMaintenanceController,
	deleteCtrl   *controllers.DeleteMaintenanceController,
) {
	group := router.Group("/maintenance")
	group.Use(security.JWTMiddleware())
	{
		group.POST("",         createCtrl.Execute)
		group.GET("",          getAllCtrl.Execute)
		group.PUT("/:id/resolve", resolveCtrl.Execute)
		group.DELETE("/:id",   deleteCtrl.Execute)
	}
}