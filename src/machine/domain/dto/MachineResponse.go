package dto

import "time"

type MachineResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Capacity  string    `json:"capacity"`
	Location  *string   `json:"location,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
