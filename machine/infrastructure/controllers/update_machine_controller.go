package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"laundry-hub-api/machine/application/usecases"
)

type UpdateMachineController struct {
	updateMachineUseCase *usecases.UpdateMachineUseCase
}

func NewUpdateMachineController(updateMachineUseCase *usecases.UpdateMachineUseCase) *UpdateMachineController {
	return &UpdateMachineController{
		updateMachineUseCase: updateMachineUseCase,
	}
}

type UpdateMachineRequest struct {
	Name     string `json:"name"`
	Capacity string `json:"capacity"`
	Location string `json:"location"`
	Status   string `json:"status"`
}

func (ctrl *UpdateMachineController) Handle(c *gin.Context) {
	id := c.Param("id")
	var req UpdateMachineRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Ejecutar use case
	machine, err := ctrl.updateMachineUseCase.Execute(
		id,
		req.Name,
		req.Capacity,
		req.Location,
		req.Status,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    machine,
	})
}