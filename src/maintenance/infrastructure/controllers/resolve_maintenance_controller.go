package controllers

import (
	"laundry-hub-api/src/maintenance/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResolveMaintenanceController struct {
	resolveMaintenance *application.ResolveMaintenance
}

func NewResolveMaintenanceController(uc *application.ResolveMaintenance) *ResolveMaintenanceController {
	return &ResolveMaintenanceController{resolveMaintenance: uc}
}

func (rc *ResolveMaintenanceController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := rc.resolveMaintenance.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registro resuelto exitosamente"})
}