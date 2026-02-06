package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"laundry-hub-api/user/application/usecases"
)

type LoginController struct {
	loginUseCase *usecases.LoginUserUseCase
}

func NewLoginController(loginUseCase *usecases.LoginUserUseCase) *LoginController {
	return &LoginController{
		loginUseCase: loginUseCase,
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (ctrl *LoginController) Handle(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Ejecutar use case
	user, token, err := ctrl.loginUseCase.Execute(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
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