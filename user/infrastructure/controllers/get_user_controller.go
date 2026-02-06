package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"laundry-hub-api/user/application/usecases"
)

type GetUserController struct {
	getUserByIDUseCase *usecases.GetUserByIDUseCase
}

func NewGetUserController(getUserByIDUseCase *usecases.GetUserByIDUseCase) *GetUserController {
	return &GetUserController{
		getUserByIDUseCase: getUserByIDUseCase,
	}
}

// GetMe obtiene el usuario actual (desde el token JWT)
func (ctrl *GetUserController) GetMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User not authenticated",
		})
		return
	}

	user, err := ctrl.getUserByIDUseCase.Execute(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
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
			"created_at": user.CreatedAt,
		},
	})
}