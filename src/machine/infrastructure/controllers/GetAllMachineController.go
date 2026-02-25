package controllers

import (
	"laundry-hub-api/src/machine/application"
	"laundry-hub-api/src/machine/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllMachinesController struct {
	getAllMachines *application.GetAllMachines
}

func NewGetAllMachinesController(getAllMachines *application.GetAllMachines) *GetAllMachinesController {
	return &GetAllMachinesController{getAllMachines: getAllMachines}
}

func (gc *GetAllMachinesController) Execute(c *gin.Context) {
	machines, err := gc.getAllMachines.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.MachineResponse
	for _, machine := range machines {
		response = append(response, dto.MachineResponse{
			ID:        machine.ID,
			Name:      machine.Name,
			Status:    machine.Status,
			Capacity:  machine.Capacity,
			Location:  machine.Location,
			CreatedAt: machine.CreatedAt,
			UpdatedAt: machine.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"machines": response})
}
