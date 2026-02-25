package dto

import "time"

type UserResponse struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	SecondName      *string   `json:"secondName,omitempty"`
	PaternalSurname string    `json:"paternalSurname"`
	MaternalSurname *string   `json:"maternalSurname,omitempty"`
	Email           string    `json:"email"`
	ImageProfile    *string   `json:"imageProfile,omitempty"`
	OAuthProvider   string    `json:"oauthProvider"`
	Role            string    `json:"role"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type LoginResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
}
