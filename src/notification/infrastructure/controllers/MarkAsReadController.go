package controllers

import (
	"laundry-hub-api/src/notification/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MarkAsReadController struct {
	markAsRead *application.MarkAsRead
}

func NewMarkAsReadController(markAsRead *application.MarkAsRead) *MarkAsReadController {
	return &MarkAsReadController{markAsRead: markAsRead}
}

func (mc *MarkAsReadController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := mc.markAsRead.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notificación marcada como leída"})
}
