package controllers

import (
	"laundry-hub-api/src/reservation/application"
	"laundry-hub-api/src/reservation/domain/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetReservationByIdController struct {
	getReservationByID *application.GetReservationByID
}

func NewGetReservationByIdController(getReservationByID *application.GetReservationByID) *GetReservationByIdController {
	return &GetReservationByIdController{getReservationByID: getReservationByID}
}

func (gc *GetReservationByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	reservation, err := gc.getReservationByID.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if reservation == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservación no encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reservation": dto.ReservationResponse{
			ID:        reservation.ID,
			UserID:    reservation.UserID,
			MachineID: reservation.MachineID,
			Status:    reservation.Status,
			StartedAt: reservation.StartedAt,
			EndedAt:   reservation.EndedAt,
		},
	})
}
