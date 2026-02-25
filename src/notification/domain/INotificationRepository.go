package domain

import "laundry-hub-api/src/notification/domain/entities"

type INotificationRepository interface {
	Save(notification *entities.Notification) (*entities.Notification, error)
	GetByID(id int) (*entities.Notification, error)
	GetByUserID(userID int) ([]*entities.Notification, error)
	MarkAsRead(id int) error
	MarkAllAsRead(userID int) error
}
