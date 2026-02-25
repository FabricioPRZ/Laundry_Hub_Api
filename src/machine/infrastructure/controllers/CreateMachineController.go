package controllers

import (
	"laundry-hub-api/src/machine/application"
	"laundry-hub-api/src/machine/domain/dto"
	"laundry-hub-api/src/machine/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMachineController struct {
	createMachine *application.CreateMachine
}

func NewCreateMachineController(createMachine *application.CreateMachine) *CreateMachineController {
	return &CreateMachineController{createMachine: createMachine}
}

func (cc *CreateMachineController) Execute(c *gin.Context) {
	var req dto.CreateMachineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	machine := &entities.Machine{
		Name:     req.Name,
		Capacity: req.Capacity,
		Location: req.Location,
	}

	saved, err := cc.createMachine.Execute(machine)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Máquina creada exitosamente",
		"machine": dto.MachineResponse{
			ID:        saved.ID,
			Name:      saved.Name,
			Status:    saved.Status,
			Capacity:  saved.Capacity,
			Location:  saved.Location,
			CreatedAt: saved.CreatedAt,
			UpdatedAt: saved.UpdatedAt,
		},
	})
}
