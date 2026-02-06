package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"laundry-hub-api/machine/application/usecases"
)

type GetMachineByIDController struct {
	getMachineByIDUseCase *usecases.GetMachineByIDUseCase
}

func NewGetMachineByIDController(getMachineByIDUseCase *usecases.GetMachineByIDUseCase) *GetMachineByIDController {
	return &GetMachineByIDController{
		getMachineByIDUseCase: getMachineByIDUseCase,
	}
}

func (ctrl *GetMachineByIDController) Handle(c *gin.Context) {
	id := c.Param("id")

	// Ejecutar use case
	machine, err := ctrl.getMachineByIDUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
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