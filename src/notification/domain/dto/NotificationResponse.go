package dto

import "time"

type NotificationResponse struct {
	ID            int       `json:"id"`
	UserID        int       `json:"userId"`
	ReservationID *int      `json:"reservationId,omitempty"`
	Message       string    `json:"message"`
	Type          string    `json:"type"`
	IsRead        bool      `json:"isRead"`
	CreatedAt     time.Time `json:"createdAt"`
}
