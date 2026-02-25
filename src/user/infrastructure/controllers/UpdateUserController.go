package controllers

import (
	"io"
	"laundry-hub-api/src/core/cloudinary"
	"laundry-hub-api/src/user/application"
	"laundry-hub-api/src/user/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	updateUser *application.UpdateUser
}

func NewUpdateUserController(updateUser *application.UpdateUser) *UpdateUserController {
	return &UpdateUserController{updateUser: updateUser}
}

func (uc *UpdateUserController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	name := c.PostForm("name")
	secondName := c.PostForm("secondName")
	paternalSurname := c.PostForm("paternalSurname")
	maternalSurname := c.PostForm("maternalSurname")
	email := c.PostForm("email")
	role := c.PostForm("role")

	if name == "" || paternalSurname == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obligatorios: name, paternalSurname, email"})
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
		ID:              id,
		Name:            name,
		SecondName:      secondNamePtr,
		PaternalSurname: paternalSurname,
		MaternalSurname: maternalSurnamePtr,
		Email:           email,
		ImageProfile:    imageProfileURL,
		Role:            role,
	}

	if err := uc.updateUser.Execute(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente"})
}
