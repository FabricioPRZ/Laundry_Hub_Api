package routes

import (
	"github.com/gin-gonic/gin"
	"laundry-hub-api/core/security"
	"laundry-hub-api/user/infrastructure/dependencies"
)

func RegisterUserRoutes(api *gin.RouterGroup, deps *dependencies.UserDependencies) {
	auth := api.Group("auth")
	{
		// Rutas públicas
		auth.POST("/register", deps.RegisterController.Handle)
		auth.POST("/login", deps.LoginController.Handle)

		// Rutas protegidas
		auth.GET("/me", security.AuthMiddleware(), deps.GetUserController.GetMe)
	}
}