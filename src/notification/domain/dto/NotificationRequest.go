package dto

type CreateNotificationRequest struct {
	UserID        int    `json:"userId" binding:"required"`
	ReservationID *int   `json:"reservationId,omitempty"`
	Message       string `json:"message" binding:"required"`
	Type          string `json:"type" binding:"required"`
}
