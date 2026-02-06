package security

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifica que el usuario esté autenticado
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener token del header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "No token provided",
			})
			c.Abort()
			return
		}

		// Extraer token (formato: "Bearer TOKEN")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Validar token
		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Guardar información del usuario en el contexto
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}

// AdminMiddleware verifica que el usuario sea administrador
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists || role != "ADMIN" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Access denied. Admin only.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}