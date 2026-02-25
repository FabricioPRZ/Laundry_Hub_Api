package controllers

import (
	"laundry-hub-api/src/user/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	deleteUser *application.DeleteUser
}

func NewDeleteUserController(deleteUser *application.DeleteUser) *DeleteUserController {
	return &DeleteUserController{deleteUser: deleteUser}
}

func (dc *DeleteUserController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := dc.deleteUser.Execute(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}
