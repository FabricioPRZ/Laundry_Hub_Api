package infrastructure

import (
	"laundry-hub-api/src/core"
	"laundry-hub-api/src/user/application"
	"laundry-hub-api/src/user/infrastructure/adapters"
	"laundry-hub-api/src/user/infrastructure/controllers"
)

type DependenciesUsers struct {
	AuthController        *controllers.AuthController
	CreateUserController  *controllers.CreateUserController
	GetAllUsersController *controllers.GetAllUsersController
	GetUserByIdController *controllers.GetUserByIdController
	UpdateUserController  *controllers.UpdateUserController
	DeleteUserController  *controllers.DeleteUserController
	OAuthController       *controllers.OAuthController
}

func InitUsers() *DependenciesUsers {
	conn := core.GetDBPool()
	userRepo := adapters.NewMySQL(conn.DB)

	authService := application.NewAuthService(userRepo)
	getAllUsers := application.NewGetAllUsers(userRepo)
	getUserById := application.NewGetUserById(userRepo)
	updateUser := application.NewUpdateUser(userRepo)
	deleteUser := application.NewDeleteUser(userRepo)
	oauthService := application.NewOAuthService(userRepo)

	return &DependenciesUsers{
		AuthController:        controllers.NewAuthController(authService),
		CreateUserController:  controllers.NewCreateUserController(authService),
		GetAllUsersController: controllers.NewGetAllUsersController(getAllUsers),
		GetUserByIdController: controllers.NewGetUserByIdController(getUserById),
		UpdateUserController:  controllers.NewUpdateUserController(updateUser),
		DeleteUserController:  controllers.NewDeleteUserController(deleteUser),
		OAuthController:       controllers.NewOAuthController(oauthService),
	}
}
