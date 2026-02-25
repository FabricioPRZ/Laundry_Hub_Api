package controllers

import (
	"laundry-hub-api/src/user/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OAuthController struct {
	oauthService *application.OAuthService
}

func NewOAuthController(oauthService *application.OAuthService) *OAuthController {
	return &OAuthController{oauthService: oauthService}
}

func (oc *OAuthController) GoogleCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Google OAuth en construcción"})
}

func (oc *OAuthController) GitHubCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GitHub OAuth en construcción"})
}
