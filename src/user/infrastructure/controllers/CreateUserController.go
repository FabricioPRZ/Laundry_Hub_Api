package controllers

import (
	"io"
	"laundry-hub-api/src/core/cloudinary"
	"laundry-hub-api/src/user/application"
	"laundry-hub-api/src/user/domain/dto"
	"laundry-hub-api/src/user/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	authService *application.AuthService
}

func NewCreateUserController(authService *application.AuthService) *CreateUserController {
	return &CreateUserController{authService: authService}
}

func (uc *CreateUserController) Execute(c *gin.Context) {
	name := c.PostForm("name")
	secondName := c.PostForm("secondName")
	paternalSurname := c.PostForm("paternalSurname")
	maternalSurname := c.PostForm("maternalSurname")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if name == "" || paternalSurname == "" || email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obligatorios: name, paternalSurname, email, password"})
		return
	}

	var imageProfileURL *string
	file, header, err := c.Request.FormFile("imageProfile")
	if err == nil {
		defer file.Close()
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar imagen"})
			return
		}
		uploadedURL, err := cloudinary.UploadAvatar(fileBytes, header.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al subir imagen"})
			return
		}
		imageProfileURL = &uploadedURL
	}

	var secondNamePtr *string
	if secondName != "" {
		secondNamePtr = &secondName
	}
	var maternalSurnamePtr *string
	if maternalSurname != "" {
		maternalSurnamePtr = &maternalSurname
	}

	user := &entities.User{
		Name:            name,
		SecondName:      secondNamePtr,
		PaternalSurname: paternalSurname,
		MaternalSurname: maternalSurnamePtr,
		Email:           email,
		Password:        &password,
		ImageProfile:    imageProfileURL,
	}

	savedUser, err := uc.authService.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado exitosamente",
		"user": dto.UserResponse{
			ID:              savedUser.ID,
			Name:            savedUser.Name,
			SecondName:      savedUser.SecondName,
			PaternalSurname: savedUser.PaternalSurname,
			MaternalSurname: savedUser.MaternalSurname,
			Email:           savedUser.Email,
			ImageProfile:    savedUser.ImageProfile,
			OAuthProvider:   savedUser.OAuthProvider,
			Role:            savedUser.Role,
			CreatedAt:       savedUser.CreatedAt,
			UpdatedAt:       savedUser.UpdatedAt,
		},
	})
}
