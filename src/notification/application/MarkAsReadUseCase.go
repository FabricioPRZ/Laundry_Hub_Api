package application

import "laundry-hub-api/src/notification/domain"

type MarkAsRead struct {
	notificationRepo domain.INotificationRepository
}

func NewMarkAsRead(notificationRepo domain.INotificationRepository) *MarkAsRead {
	return &MarkAsRead{notificationRepo: notificationRepo}
}

func (mr *MarkAsRead) Execute(id int) error {
	return mr.notificationRepo.MarkAsRead(id)
}
