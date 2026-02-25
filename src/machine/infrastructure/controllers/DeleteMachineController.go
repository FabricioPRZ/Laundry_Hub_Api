package controllers

import (
	"laundry-hub-api/src/machine/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteMachineController struct {
	deleteMachine *application.DeleteMachine
}

func NewDeleteMachineController(deleteMachine *application.DeleteMachine) *DeleteMachineController {
	return &DeleteMachineController{deleteMachine: deleteMachine}
}

func (dc *DeleteMachineController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := dc.deleteMachine.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Máquina eliminada exitosamente"})
}
