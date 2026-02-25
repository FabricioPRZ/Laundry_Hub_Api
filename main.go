package main

import (
	"laundry-hub-api/src/core/cloudinary"
	ws "laundry-hub-api/src/core/websocket"
	machineInfrastructure "laundry-hub-api/src/machine/infrastructure"
	machineRoutes "laundry-hub-api/src/machine/infrastructure/routes"
	notificationInfrastructure "laundry-hub-api/src/notification/infrastructure"
	notificationRoutes "laundry-hub-api/src/notification/infrastructure/routes"
	reservationInfrastructure "laundry-hub-api/src/reservation/infrastructure"
	reservationRoutes "laundry-hub-api/src/reservation/infrastructure/routes"
	userInfrastructure "laundry-hub-api/src/user/infrastructure"
	userRoutes "laundry-hub-api/src/user/infrastructure/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"laundry-hub-api/src/core/security"
)

func main() {
	cloudinary.InitCloudinary()
	ws.InitWebSocket()

	userDeps := userInfrastructure.InitUsers()
	machineDeps := machineInfrastructure.InitMachines()
	reservationDeps := reservationInfrastructure.InitReservations()
	notificationDeps := notificationInfrastructure.InitNotifications()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/ws", security.JWTMiddleware(), ws.HandleConnection)

	userRoutes.ConfigureUserRoutes(
		router,
		userDeps.AuthController,
		userDeps.CreateUserController,
		userDeps.GetAllUsersController,
		userDeps.GetUserByIdController,
		userDeps.UpdateUserController,
		userDeps.DeleteUserController,
		userDeps.OAuthController,
	)

	machineRoutes.ConfigureMachineRoutes(
		router,
		machineDeps.CreateMachineController,
		machineDeps.GetAllMachinesController,
		machineDeps.GetMachineByIdController,
		machineDeps.UpdateMachineController,
		machineDeps.DeleteMachineController,
	)

	reservationRoutes.ConfigureReservationRoutes(
		router,
		reservationDeps.CreateReservationController,
		reservationDeps.CancelReservationController,
		reservationDeps.CompleteReservationController,
		reservationDeps.GetReservationByIdController,
		reservationDeps.GetReservationsByUserController,
	)

	notificationRoutes.ConfigureNotificationRoutes(
		router,
		notificationDeps.CreateNotificationController,
		notificationDeps.GetNotificationsByUserController,
		notificationDeps.MarkAsReadController,
		notificationDeps.MarkAllAsReadController,
	)

	log.Println("Servidor corriendo en http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
