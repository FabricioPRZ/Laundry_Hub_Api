package routes

import (
	"laundry-hub-api/src/core/security"
	"laundry-hub-api/src/user/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureUserRoutes(
	router *gin.Engine,
	authCtrl *controllers.AuthController,
	createUserCtrl *controllers.CreateUserController,
	getAllUsersCtrl *controllers.GetAllUsersController,
	getUserByIdCtrl *controllers.GetUserByIdController,
	updateUserCtrl *controllers.UpdateUserController,
	deleteUserCtrl *controllers.DeleteUserController,
	oauthCtrl *controllers.OAuthController,
) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authCtrl.Login)
		authGroup.POST("/register", createUserCtrl.Execute)
		authGroup.POST("/logout", authCtrl.Logout)
		authGroup.POST("/refresh", authCtrl.RefreshToken)
		authGroup.GET("/profile", security.JWTMiddleware(), authCtrl.GetProfile)
		authGroup.GET("/verify", security.JWTMiddleware(), authCtrl.VerifyToken)
		authGroup.GET("/google/callback", oauthCtrl.GoogleCallback)
		authGroup.GET("/github/callback", oauthCtrl.GitHubCallback)
	}

	userGroup := router.Group("/users")
	userGroup.Use(security.JWTMiddleware())
	{
		userGroup.GET("", getAllUsersCtrl.Execute)
		userGroup.GET("/:id", getUserByIdCtrl.Execute)
		userGroup.PUT("/:id", updateUserCtrl.Execute)
		userGroup.DELETE("/:id", deleteUserCtrl.Execute)
	}
}
