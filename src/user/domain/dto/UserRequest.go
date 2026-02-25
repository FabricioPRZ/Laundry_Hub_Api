package dto

type RegisterRequest struct {
	Name            string  `json:"name" binding:"required"`
	SecondName      *string `json:"secondName,omitempty"`
	PaternalSurname string  `json:"paternalSurname" binding:"required"`
	MaternalSurname *string `json:"maternalSurname,omitempty"`
	Email           string  `json:"email" binding:"required,email"`
	Password        string  `json:"password" binding:"required,min=6"`
	ImageProfile    *string `json:"imageProfile,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Name            string  `json:"name" binding:"required"`
	SecondName      *string `json:"secondName,omitempty"`
	PaternalSurname string  `json:"paternalSurname" binding:"required"`
	MaternalSurname *string `json:"maternalSurname,omitempty"`
	Email           string  `json:"email" binding:"required,email"`
	ImageProfile    *string `json:"imageProfile,omitempty"`
}
