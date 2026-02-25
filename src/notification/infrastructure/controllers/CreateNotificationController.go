package controllers

import (
	"laundry-hub-api/src/notification/application"
	"laundry-hub-api/src/notification/domain/dto"
	"laundry-hub-api/src/notification/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateNotificationController struct {
	createNotification *application.CreateNotification
}

func NewCreateNotificationController(createNotification *application.CreateNotification) *CreateNotificationController {
	return &CreateNotificationController{createNotification: createNotification}
}

func (cc *CreateNotificationController) Execute(c *gin.Context) {
	var req dto.CreateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	notification := &entities.Notification{
		UserID:        req.UserID,
		ReservationID: req.ReservationID,
		Message:       req.Message,
		Type:          req.Type,
	}

	saved, err := cc.createNotification.Execute(notification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Notificación creada exitosamente",
		"notification": dto.NotificationResponse{
			ID:            saved.ID,
			UserID:        saved.UserID,
			ReservationID: saved.ReservationID,
			Message:       saved.Message,
			Type:          saved.Type,
			IsRead:        saved.IsRead,
			CreatedAt:     saved.CreatedAt,
		},
	})
}
