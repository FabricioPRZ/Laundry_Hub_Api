package entities

import "time"

type MaintenanceRecord struct {
	ID          int        `json:"id"`
	MachineID   int        `json:"machine_id"`
	MachineName string     `json:"machine_name"`
	Description string     `json:"description"`
	IsResolved  bool       `json:"is_resolved"`
	ResolvedAt  *time.Time `json:"resolved_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}