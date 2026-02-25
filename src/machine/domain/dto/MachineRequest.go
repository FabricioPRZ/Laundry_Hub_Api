package dto

type CreateMachineRequest struct {
	Name     string  `json:"name" binding:"required"`
	Capacity string  `json:"capacity" binding:"required"`
	Location *string `json:"location,omitempty"`
}

type UpdateMachineRequest struct {
	Name     string  `json:"name" binding:"required"`
	Status   string  `json:"status" binding:"required"`
	Capacity string  `json:"capacity" binding:"required"`
	Location *string `json:"location,omitempty"`
}
