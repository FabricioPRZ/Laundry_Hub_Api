package infrastructure

import (
	"laundry-hub-api/src/core"
	machineAdapters "laundry-hub-api/src/machine/infrastructure/adapters"
	notificationAdapters "laundry-hub-api/src/notification/infrastructure/adapters"
	"laundry-hub-api/src/reservation/application"
	"laundry-hub-api/src/reservation/infrastructure/adapters"
	"laundry-hub-api/src/reservation/infrastructure/controllers"
)

type DependenciesReservations struct {
	CreateReservationController     *controllers.CreateReservationController
	CancelReservationController     *controllers.CancelReservationController
	CompleteReservationController   *controllers.CompleteReservationController
	GetReservationByIdController    *controllers.GetReservationByIdController
	GetReservationsByUserController *controllers.GetReservationsByUserController
}

func InitReservations() *DependenciesReservations {
	conn := core.GetDBPool()
	reservationRepo := adapters.NewMySQL(conn.DB)
	machineRepo := machineAdapters.NewMySQL(conn.DB)
	notificationRepo := notificationAdapters.NewMySQL(conn.DB)

	createReservation := application.NewCreateReservation(reservationRepo, machineRepo, notificationRepo)
	cancelReservation := application.NewCancelReservation(reservationRepo, machineRepo)
	completeReservation := application.NewCompleteReservation(reservationRepo, machineRepo)
	getReservationByID := application.NewGetReservationByID(reservationRepo)
	getReservationsByUser := application.NewGetReservationsByUser(reservationRepo)

	return &DependenciesReservations{
		CreateReservationController:     controllers.NewCreateReservationController(createReservation),
		CancelReservationController:     controllers.NewCancelReservationController(cancelReservation),
		CompleteReservationController:   controllers.NewCompleteReservationController(completeReservation),
		GetReservationByIdController:    controllers.NewGetReservationByIdController(getReservationByID),
		GetReservationsByUserController: controllers.NewGetReservationsByUserController(getReservationsByUser),
	}
}
