package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"laundry-hub-api/machine/application/usecases"
)

type GetMachinesController struct {
	getMachinesUseCase *usecases.GetMachinesUseCase
}

func NewGetMachinesController(getMachinesUseCase *usecases.GetMachinesUseCase) *GetMachinesController {
	return &GetMachinesController{
		getMachinesUseCase: getMachinesUseCase,
	}
}

func (ctrl *GetMachinesController) Handle(c *gin.Context) {
	// Ejecutar use case
	machines, err := ctrl.getMachinesUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    machines,
	})
}