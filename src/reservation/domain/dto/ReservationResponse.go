package dto

import "time"

type ReservationResponse struct {
	ID        int        `json:"id"`
	UserID    int        `json:"userId"`
	MachineID int        `json:"machineId"`
	Status    string     `json:"status"`
	StartedAt time.Time  `json:"startedAt"`
	EndedAt   *time.Time `json:"endedAt,omitempty"`
}
