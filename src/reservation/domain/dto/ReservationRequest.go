package dto

type CreateReservationRequest struct {
	MachineID int `json:"machineId" binding:"required"`
}
