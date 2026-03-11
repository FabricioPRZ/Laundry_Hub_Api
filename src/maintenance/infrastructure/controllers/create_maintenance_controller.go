package controllers

import (
	"laundry-hub-api/src/maintenance/application"
	"laundry-hub-api/src/maintenance/domain/dto"
	"laundry-hub-api/src/maintenance/infrastructure/adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMaintenanceController struct {
	createMaintenance *application.CreateMaintenance
}

func NewCreateMaintenanceController(uc *application.CreateMaintenance) *CreateMaintenanceController {
	return &CreateMaintenanceController{createMaintenance: uc}
}

func (cc *CreateMaintenanceController) Execute(c *gin.Context) {
	var req dto.CreateMaintenanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	userID, _ := c.Get("user_id")

	record, err := cc.createMaintenance.Execute(userID.(int), req.MachineID, req.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registro de mantenimiento creado",
		"record": dto.MaintenanceResponse{
			ID:          record.ID,
			MachineID:   record.MachineID,
			MachineName: record.MachineName,
			Description: record.Description,
			IsResolved:  record.IsResolved,
			ResolvedAt:  record.ResolvedAt,
			StartDate:   record.CreatedAt.Format("2 Jan 2006"),
			DaysElapsed: adapters.DaysElapsed(record.CreatedAt),
		},
	})
}