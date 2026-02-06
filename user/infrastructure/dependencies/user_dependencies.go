package dependencies

import (
	"database/sql"

	"laundry-hub-api/user/application/usecases"
	"laundry-hub-api/user/infrastructure/adapters"
	"laundry-hub-api/user/infrastructure/controllers"
)

type UserDependencies struct {
	RegisterController *controllers.RegisterController
	LoginController    *controllers.LoginController
	GetUserController  *controllers.GetUserController
}

func NewUserDependencies(db *sql.DB) *UserDependencies {
	// Adapter
	userRepo := adapters.NewUserMySQLRepository(db)

	// Use Cases
	registerUseCase := usecases.NewRegisterUserUseCase(userRepo)
	loginUseCase := usecases.NewLoginUserUseCase(userRepo)
	getUserByIDUseCase := usecases.NewGetUserByIDUseCase(userRepo)

	// Controllers
	registerController := controllers.NewRegisterController(registerUseCase)
	loginController := controllers.NewLoginController(loginUseCase)
	getUserController := controllers.NewGetUserController(getUserByIDUseCase)

	return &UserDependencies{
		RegisterController: registerController,
		LoginController:    loginController,
		GetUserController:  getUserController,
	}
}