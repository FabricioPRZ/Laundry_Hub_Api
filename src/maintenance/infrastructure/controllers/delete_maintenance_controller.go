package controllers

import (
	"laundry-hub-api/src/maintenance/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteMaintenanceController struct {
	deleteMaintenance *application.DeleteMaintenance
}

func NewDeleteMaintenanceController(uc *application.DeleteMaintenance) *DeleteMaintenanceController {
	return &DeleteMaintenanceController{deleteMaintenance: uc}
}

func (dc *DeleteMaintenanceController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := dc.deleteMaintenance.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registro eliminado exitosamente"})
}