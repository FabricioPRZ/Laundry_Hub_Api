package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"laundry-hub-api/machine/application/usecases"
)

type DeleteMachineController struct {
	deleteMachineUseCase *usecases.DeleteMachineUseCase
}

func NewDeleteMachineController(deleteMachineUseCase *usecases.DeleteMachineUseCase) *DeleteMachineController {
	return &DeleteMachineController{
		deleteMachineUseCase: deleteMachineUseCase,
	}
}

func (ctrl *DeleteMachineController) Handle(c *gin.Context) {
	id := c.Param("id")

	// Ejecutar use case
	err := ctrl.deleteMachineUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Machine deleted successfully",
	})
}