package controllers

import (
	"laundry-hub-api/src/core/security"
	"laundry-hub-api/src/user/application"
	"laundry-hub-api/src/user/domain/dto"
	"laundry-hub-api/src/user/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *application.AuthService
}

func NewAuthController(authService *application.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	user, err := ac.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := security.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	refreshToken, err := security.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar refresh token"})
		return
	}

	security.SetAuthCookie(c.Writer, accessToken)
	security.SetRefreshCookie(c.Writer, refreshToken)

	c.JSON(http.StatusOK, dto.LoginResponse{
		Message: "Login exitoso",
		User: dto.UserResponse{
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

func (ac *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	user := &entities.User{
		Name:            req.Name,
		SecondName:      req.SecondName,
		PaternalSurname: req.PaternalSurname,
		MaternalSurname: req.MaternalSurname,
		Email:           req.Email,
		Password:        &req.Password,
		ImageProfile:    req.ImageProfile,
	}

	savedUser, err := ac.authService.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
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

func (ac *AuthController) Logout(c *gin.Context) {
	security.ClearAuthCookies(c.Writer)
	c.JSON(http.StatusOK, gin.H{"message": "Logout exitoso"})
}

func (ac *AuthController) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token no encontrado"})
		return
	}

	claims, err := security.ValidateRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido"})
		return
	}

	user, err := ac.authService.GetUserByID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	newAccessToken, err := security.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	security.SetAuthCookie(c.Writer, newAccessToken)
	c.JSON(http.StatusOK, gin.H{"message": "Token renovado"})
}

func (ac *AuthController) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autenticado"})
		return
	}

	user, err := ac.authService.GetUserByID(userID.(int))
	if err != nil {
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

func (ac *AuthController) VerifyToken(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")
	role, _ := c.Get("role")

	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"user": gin.H{
			"id":    userID,
			"email": email,
			"role":  role,
		},
	})
}
