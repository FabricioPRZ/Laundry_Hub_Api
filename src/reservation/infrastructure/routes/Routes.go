package routes

import (
	"laundry-hub-api/src/core/security"
	"laundry-hub-api/src/reservation/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureReservationRoutes(
	router *gin.Engine,
	createCtrl *controllers.CreateReservationController,
	cancelCtrl *controllers.CancelReservationController,
	completeCtrl *controllers.CompleteReservationController,
	getByIDCtrl *controllers.GetReservationByIdController,
	getByUserCtrl *controllers.GetReservationsByUserController,
) {
	reservationGroup := router.Group("/reservations")
	reservationGroup.Use(security.JWTMiddleware())
	{
		reservationGroup.POST("", createCtrl.Execute)
		reservationGroup.GET("/:id", getByIDCtrl.Execute)
		reservationGroup.GET("/my", getByUserCtrl.Execute)
		reservationGroup.PUT("/:id/cancel", cancelCtrl.Execute)
		reservationGroup.PUT("/:id/complete", completeCtrl.Execute)
	}
}
