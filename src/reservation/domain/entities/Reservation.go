package entities

import "time"

type Reservation struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	MachineID int        `json:"machine_id"`
	Status    string     `json:"status"`
	StartedAt time.Time  `json:"started_at"`
	EndedAt   *time.Time `json:"ended_at,omitempty"`
}
