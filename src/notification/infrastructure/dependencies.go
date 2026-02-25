package infrastructure

import (
	"laundry-hub-api/src/core"
	"laundry-hub-api/src/notification/application"
	"laundry-hub-api/src/notification/infrastructure/adapters"
	"laundry-hub-api/src/notification/infrastructure/controllers"
)

type DependenciesNotifications struct {
	CreateNotificationController     *controllers.CreateNotificationController
	GetNotificationsByUserController *controllers.GetNotificationsByUserController
	MarkAsReadController             *controllers.MarkAsReadController
	MarkAllAsReadController          *controllers.MarkAllAsReadController
}

func InitNotifications() *DependenciesNotifications {
	conn := core.GetDBPool()
	notificationRepo := adapters.NewMySQL(conn.DB)

	createNotification := application.NewCreateNotification(notificationRepo)
	getNotificationsByUser := application.NewGetNotificationsByUser(notificationRepo)
	markAsRead := application.NewMarkAsRead(notificationRepo)
	markAllAsRead := application.NewMarkAllAsRead(notificationRepo)

	return &DependenciesNotifications{
		CreateNotificationController:     controllers.NewCreateNotificationController(createNotification),
		GetNotificationsByUserController: controllers.NewGetNotificationsByUserController(getNotificationsByUser),
		MarkAsReadController:             controllers.NewMarkAsReadController(markAsRead),
		MarkAllAsReadController:          controllers.NewMarkAllAsReadController(markAllAsRead),
	}
}
