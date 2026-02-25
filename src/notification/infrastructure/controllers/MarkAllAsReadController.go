package controllers

import (
	"laundry-hub-api/src/notification/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MarkAllAsReadController struct {
	markAllAsRead *application.MarkAllAsRead
}

func NewMarkAllAsReadController(markAllAsRead *application.MarkAllAsRead) *MarkAllAsReadController {
	return &MarkAllAsReadController{markAllAsRead: markAllAsRead}
}

func (mc *MarkAllAsReadController) Execute(c *gin.Context) {
	userID, _ := c.Get("user_id")

	if err := mc.markAllAsRead.Execute(userID.(int)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todas las notificaciones marcadas como leídas"})
}
