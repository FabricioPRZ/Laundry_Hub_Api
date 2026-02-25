package controllers

import (
	"laundry-hub-api/src/reservation/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CancelReservationController struct {
	cancelReservation *application.CancelReservation
}

func NewCancelReservationController(cancelReservation *application.CancelReservation) *CancelReservationController {
	return &CancelReservationController{cancelReservation: cancelReservation}
}

func (cc *CancelReservationController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userID, _ := c.Get("user_id")

	if err := cc.cancelReservation.Execute(id, userID.(int)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservación cancelada exitosamente"})
}
