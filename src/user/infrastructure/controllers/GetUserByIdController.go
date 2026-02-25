package controllers

import (
	"laundry-hub-api/src/user/application"
	"laundry-hub-api/src/user/domain/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUserByIdController struct {
	getUserById *application.GetUserById
}

func NewGetUserByIdController(getUserById *application.GetUserById) *GetUserByIdController {
	return &GetUserByIdController{getUserById: getUserById}
}

func (gc *GetUserByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := gc.getUserById.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": dto.UserResponse{
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
		},
	})
}
