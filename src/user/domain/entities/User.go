package entities

import "time"

type User struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	SecondName      *string   `json:"second_name,omitempty"`
	PaternalSurname string    `json:"paternal_surname"`
	MaternalSurname *string   `json:"maternal_surname,omitempty"`
	Email           string    `json:"email"`
	Password        *string   `json:"password,omitempty"`
	ImageProfile    *string   `json:"image_profile,omitempty"`
	OAuthProvider   string    `json:"oauth_provider"`
	OAuthID         *string   `json:"oauth_id,omitempty"`
	Role            string    `json:"role"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
