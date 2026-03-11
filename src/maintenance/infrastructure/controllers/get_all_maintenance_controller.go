package controllers

import (
	"laundry-hub-api/src/maintenance/application"
	"laundry-hub-api/src/maintenance/domain/dto"
	"laundry-hub-api/src/maintenance/infrastructure/adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllMaintenanceController struct {
	getAllMaintenance *application.GetAllMaintenance
}

func NewGetAllMaintenanceController(uc *application.GetAllMaintenance) *GetAllMaintenanceController {
	return &GetAllMaintenanceController{getAllMaintenance: uc}
}

func (gc *GetAllMaintenanceController) Execute(c *gin.Context) {
	records, err := gc.getAllMaintenance.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]dto.MaintenanceResponse, 0)
	for _, r := range records {
		response = append(response, dto.MaintenanceResponse{
			ID:          r.ID,
			MachineID:   r.MachineID,
			MachineName: r.MachineName,
			Description: r.Description,
			IsResolved:  r.IsResolved,
			ResolvedAt:  r.ResolvedAt,
			StartDate:   r.CreatedAt.Format("2 Jan 2006"),
			DaysElapsed: adapters.DaysElapsed(r.CreatedAt),
		})
	}

	c.JSON(http.StatusOK, gin.H{"records": response})
}