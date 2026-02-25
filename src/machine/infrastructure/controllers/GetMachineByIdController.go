package controllers

import (
	"laundry-hub-api/src/machine/application"
	"laundry-hub-api/src/machine/domain/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetMachineByIdController struct {
	getMachineByID *application.GetMachineByID
}

func NewGetMachineByIdController(getMachineByID *application.GetMachineByID) *GetMachineByIdController {
	return &GetMachineByIdController{getMachineByID: getMachineByID}
}

func (gc *GetMachineByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	machine, err := gc.getMachineByID.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if machine == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Máquina no encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"machine": dto.MachineResponse{
			ID:        machine.ID,
			Name:      machine.Name,
			Status:    machine.Status,
			Capacity:  machine.Capacity,
			Location:  machine.Location,
			CreatedAt: machine.CreatedAt,
			UpdatedAt: machine.UpdatedAt,
		},
	})
}
