package controllers

import (
	"laundry-hub-api/src/reservation/application"
	"laundry-hub-api/src/reservation/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateReservationController struct {
	createReservation *application.CreateReservation
}

func NewCreateReservationController(createReservation *application.CreateReservation) *CreateReservationController {
	return &CreateReservationController{createReservation: createReservation}
}

func (cc *CreateReservationController) Execute(c *gin.Context) {
	var req dto.CreateReservationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	userID, _ := c.Get("user_id")

	saved, err := cc.createReservation.Execute(userID.(int), req.MachineID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Reservación creada exitosamente",
		"reservation": dto.ReservationResponse{
			ID:        saved.ID,
			UserID:    saved.UserID,
			MachineID: saved.MachineID,
			Status:    saved.Status,
			StartedAt: saved.StartedAt,
			EndedAt:   saved.EndedAt,
		},
	})
}
