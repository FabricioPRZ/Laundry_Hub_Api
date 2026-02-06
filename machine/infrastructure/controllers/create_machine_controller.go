package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"laundry-hub-api/machine/application/usecases"
)

type CreateMachineController struct {
	createMachineUseCase *usecases.CreateMachineUseCase
}

func NewCreateMachineController(createMachineUseCase *usecases.CreateMachineUseCase) *CreateMachineController {
	return &CreateMachineController{
		createMachineUseCase: createMachineUseCase,
	}
}

type CreateMachineRequest struct {
	Name     string `json:"name" binding:"required"`
	Capacity string `json:"capacity" binding:"required"`
	Location string `json:"location" binding:"required"`
	Status   string `json:"status"`
}

func (ctrl *CreateMachineController) Handle(c *gin.Context) {
	var req CreateMachineRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Si no se proporciona status, usar AVAILABLE por defecto
	if req.Status == "" {
		req.Status = "AVAILABLE"
	}

	// Ejecutar use case
	machine, err := ctrl.createMachineUseCase.Execute(
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

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    machine,
	})
}