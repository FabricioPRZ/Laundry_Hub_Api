package routes

import (
	"laundry-hub-api/src/core/security"
	"laundry-hub-api/src/notification/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureNotificationRoutes(
	router *gin.Engine,
	createCtrl *controllers.CreateNotificationController,
	getByUserCtrl *controllers.GetNotificationsByUserController,
	markAsReadCtrl *controllers.MarkAsReadController,
	markAllAsReadCtrl *controllers.MarkAllAsReadController,
) {
	notificationGroup := router.Group("/notifications")
	notificationGroup.Use(security.JWTMiddleware())
	{
		notificationGroup.POST("", createCtrl.Execute)
		notificationGroup.GET("/my", getByUserCtrl.Execute)
		notificationGroup.PUT("/:id/read", markAsReadCtrl.Execute)
		notificationGroup.PUT("/read-all", markAllAsReadCtrl.Execute)
	}
}
