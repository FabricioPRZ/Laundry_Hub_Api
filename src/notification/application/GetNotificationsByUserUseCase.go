package application

import (
	"laundry-hub-api/src/notification/domain"
	"laundry-hub-api/src/notification/domain/entities"
)

type GetNotificationsByUser struct {
	notificationRepo domain.INotificationRepository
}

func NewGetNotificationsByUser(notificationRepo domain.INotificationRepository) *GetNotificationsByUser {
	return &GetNotificationsByUser{notificationRepo: notificationRepo}
}

func (gn *GetNotificationsByUser) Execute(userID int) ([]*entities.Notification, error) {
	return gn.notificationRepo.GetByUserID(userID)
}
