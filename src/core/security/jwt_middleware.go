package security

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No autenticado - token no encontrado",
			})
			c.Abort()
			return
		}

		if cookie == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token vacío",
			})
			c.Abort()
			return
		}

		claims, err := ValidateJWT(cookie)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token inválido o expirado",
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Rol no encontrado en el token",
			})
			c.Abort()
			return
		}

		userRole, ok := role.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error al procesar el rol",
			})
			c.Abort()
			return
		}

		if userRole != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "No tienes permisos para acceder a este recurso",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func RequireAnyRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Rol no encontrado en el token",
			})
			c.Abort()
			return
		}

		userRole, ok := role.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error al procesar el rol",
			})
			c.Abort()
			return
		}

		hasPermission := false
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "No tienes permisos para acceder a este recurso",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func OptionalJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("access_token")
		if err != nil || cookie == "" {
			c.Next()
			return
		}

		claims, err := ValidateJWT(cookie)
		if err != nil {
			c.Next()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}
