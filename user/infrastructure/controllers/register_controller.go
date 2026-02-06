package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"laundry-hub-api/core/security"
	"laundry-hub-api/user/application/usecases"
)

type RegisterController struct {
	registerUseCase *usecases.RegisterUserUseCase
}

func NewRegisterController(registerUseCase *usecases.RegisterUserUseCase) *RegisterController {
	return &RegisterController{
		registerUseCase: registerUseCase,
	}
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (ctrl *RegisterController) Handle(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Ejecutar use case
	user, err := ctrl.registerUseCase.Execute(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Generar token
	token, err := security.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error generating token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"role":       user.Role,
			"token":      token,
			"created_at": user.CreatedAt,
		},
	})
}