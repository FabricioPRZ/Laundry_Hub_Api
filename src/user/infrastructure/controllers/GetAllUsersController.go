package controllers

import (
	"laundry-hub-api/src/user/application"
	"laundry-hub-api/src/user/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllUsersController struct {
	getAllUsers *application.GetAllUsers
}

func NewGetAllUsersController(getAllUsers *application.GetAllUsers) *GetAllUsersController {
	return &GetAllUsersController{getAllUsers: getAllUsers}
}

func (gc *GetAllUsersController) Execute(c *gin.Context) {
	users, err := gc.getAllUsers.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:              user.ID,
			Name:            user.Name,
			SecondName:      user.SecondName,
			PaternalSurname: user.PaternalSurname,
			MaternalSurname: user.MaternalSurname,
			Email:           user.Email,
			ImageProfile:    user.ImageProfile,
			OAuthProvider:   user.OAuthProvider,
			Role:            user.Role,
			CreatedAt:       user.CreatedAt,
			UpdatedAt:       user.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"users": userResponses})
}
