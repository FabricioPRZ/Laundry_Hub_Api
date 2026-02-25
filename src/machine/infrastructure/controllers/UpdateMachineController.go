package controllers

import (
	"laundry-hub-api/src/machine/application"
	"laundry-hub-api/src/machine/domain/dto"
	"laundry-hub-api/src/machine/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateMachineController struct {
	updateMachine *application.UpdateMachine
}

func NewUpdateMachineController(updateMachine *application.UpdateMachine) *UpdateMachineController {
	return &UpdateMachineController{updateMachine: updateMachine}
}

func (uc *UpdateMachineController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.UpdateMachineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	machine := &entities.Machine{
		ID:       id,
		Name:     req.Name,
		Status:   req.Status,
		Capacity: req.Capacity,
		Location: req.Location,
	}

	if err := uc.updateMachine.Execute(machine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Máquina actualizada exitosamente"})
}
