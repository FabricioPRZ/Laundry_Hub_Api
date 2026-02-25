package controllers

import (
	"laundry-hub-api/src/notification/application"
	"laundry-hub-api/src/notification/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetNotificationsByUserController struct {
	getNotificationsByUser *application.GetNotificationsByUser
}

func NewGetNotificationsByUserController(getNotificationsByUser *application.GetNotificationsByUser) *GetNotificationsByUserController {
	return &GetNotificationsByUserController{getNotificationsByUser: getNotificationsByUser}
}

func (gc *GetNotificationsByUserController) Execute(c *gin.Context) {
	userID, _ := c.Get("user_id")

	notifications, err := gc.getNotificationsByUser.Execute(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.NotificationResponse
	for _, n := range notifications {
		response = append(response, dto.NotificationResponse{
			ID:            n.ID,
			UserID:        n.UserID,
			ReservationID: n.ReservationID,
			Message:       n.Message,
			Type:          n.Type,
			IsRead:        n.IsRead,
			CreatedAt:     n.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"notifications": response})
}
