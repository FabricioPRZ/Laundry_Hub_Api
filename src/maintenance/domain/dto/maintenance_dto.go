package dto

import "time"

type CreateMaintenanceRequest struct {
	MachineID   int    `json:"machineId"   binding:"required"`
	Description string `json:"description" binding:"required"`
}

type MaintenanceResponse struct {
	ID          int        `json:"id"`
	MachineID   int        `json:"machineId"`
	MachineName string     `json:"machineName"`
	Description string     `json:"description"`
	IsResolved  bool       `json:"isResolved"`
	ResolvedAt  *time.Time `json:"resolvedAt,omitempty"`
	StartDate   string     `json:"startDate"`
	DaysElapsed int        `json:"daysElapsed"`
}