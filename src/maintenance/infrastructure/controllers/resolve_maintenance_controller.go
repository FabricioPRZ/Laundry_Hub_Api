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

	userID, _ := c.Get("user_id")

	if err := rc.resolveMaintenance.Execute(id, userID.(int)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registro resuelto exitosamente"})
}