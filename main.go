package main

import (
	"laundry-hub-api/src/core/cloudinary"
	machineInfrastructure "laundry-hub-api/src/machine/infrastructure"
	machineRoutes "laundry-hub-api/src/machine/infrastructure/routes"
	userInfrastructure "laundry-hub-api/src/user/infrastructure"
	userRoutes "laundry-hub-api/src/user/infrastructure/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cloudinary.InitCloudinary()

	userDeps := userInfrastructure.InitUsers()
	machineDeps := machineInfrastructure.InitMachines()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

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

	log.Println("Servidor corriendo en http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
