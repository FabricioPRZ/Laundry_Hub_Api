package controllers

import (
	"laundry-hub-api/src/reservation/application"
	"laundry-hub-api/src/reservation/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetReservationsByUserController struct {
	getReservationsByUser *application.GetReservationsByUser
}

func NewGetReservationsByUserController(getReservationsByUser *application.GetReservationsByUser) *GetReservationsByUserController {
	return &GetReservationsByUserController{getReservationsByUser: getReservationsByUser}
}

func (gc *GetReservationsByUserController) Execute(c *gin.Context) {
	userID, _ := c.Get("user_id")

	reservations, err := gc.getReservationsByUser.Execute(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.ReservationResponse
	for _, r := range reservations {
		response = append(response, dto.ReservationResponse{
			ID:        r.ID,
			UserID:    r.UserID,
			MachineID: r.MachineID,
			Status:    r.Status,
			StartedAt: r.StartedAt,
			EndedAt:   r.EndedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"reservations": response})
}
