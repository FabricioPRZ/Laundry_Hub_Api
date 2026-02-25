package application

import "laundry-hub-api/src/notification/domain"

type MarkAllAsRead struct {
	notificationRepo domain.INotificationRepository
}

func NewMarkAllAsRead(notificationRepo domain.INotificationRepository) *MarkAllAsRead {
	return &MarkAllAsRead{notificationRepo: notificationRepo}
}

func (mr *MarkAllAsRead) Execute(userID int) error {
	return mr.notificationRepo.MarkAllAsRead(userID)
}
