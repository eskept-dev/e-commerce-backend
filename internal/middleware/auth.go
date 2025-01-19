package middleware

import (
	"eskept/internal/app/context"
	jwt "eskept/internal/utils/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware extracts email and role from JWT token and sets them in the 	context
func AuthMiddleware(appCtx *context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is required"})
			return
		}

		// Extract the token from the Authorization header
		// Format: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header format"})
			return
		}

		tokenString := parts[1]
		claims, err := jwt.ValidateToken(tokenString, appCtx)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Set email and role in the context for use in handlers
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}
