package entities

import "time"

type Notification struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	ReservationID *int      `json:"reservation_id,omitempty"`
	Message       string    `json:"message"`
	Type          string    `json:"type"`
	IsRead        bool      `json:"is_read"`
	CreatedAt     time.Time `json:"created_at"`
}
