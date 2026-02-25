package application

import (
	"errors"
	"laundry-hub-api/src/notification/domain"
	"laundry-hub-api/src/notification/domain/entities"
)

type CreateNotification struct {
	notificationRepo domain.INotificationRepository
}

func NewCreateNotification(notificationRepo domain.INotificationRepository) *CreateNotification {
	return &CreateNotification{notificationRepo: notificationRepo}
}

func (cn *CreateNotification) Execute(notification *entities.Notification) (*entities.Notification, error) {
	if notification.Message == "" {
		return nil, errors.New("el mensaje es obligatorio")
	}
	if notification.Type == "" {
		return nil, errors.New("el tipo es obligatorio")
	}

	return cn.notificationRepo.Save(notification)
}
