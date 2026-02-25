package controllers

import (
	"context"
	"laundry-hub-api/src/core/security"
	"laundry-hub-api/src/user/application"
	"laundry-hub-api/src/user/domain/dto"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

type OAuthController struct {
	oauthService *application.OAuthService
}

func NewOAuthController(oauthService *application.OAuthService) *OAuthController {
	return &OAuthController{oauthService: oauthService}
}

func (oc *OAuthController) GoogleMobile(c *gin.Context) {
	var req struct {
		IDToken string `json:"idToken" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "idToken es requerido"})
		return
	}

	clientID := os.Getenv("GOOGLE_WEB_CLIENT_ID")

	payload, err := idtoken.Validate(context.Background(), req.IDToken, clientID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de Google inválido"})
		return
	}

	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)
	googleID := payload.Subject

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo obtener el email de Google"})
		return
	}

	user, err := oc.oauthService.FindOrCreateOAuthUser(email, "GOOGLE", googleID, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar usuario de Google"})
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Login con Google exitoso",
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

func (oc *OAuthController) GoogleCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Google OAuth web en construcción"})
}

func (oc *OAuthController) GitHubCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GitHub OAuth en construcción"})
}
