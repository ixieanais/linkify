// Package middlewares
package middlewares

import (
	"net/http"
	"strings"

	"linkify/internal/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokens []string

		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" || strings.TrimSpace(parts[1]) == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"detail": "invalid authorization format. Use 'Bearer TOKEN'",
				})
				return
			}

			tokens = append(tokens, strings.TrimSpace(parts[1]))
		}

		if session, err := c.Cookie("session"); err == nil && session != "" {
			tokens = append(tokens, session)
		}

		if len(tokens) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"detail": "authorization header or session cookie is required",
			})
			return
		}

		for _, tokenString := range tokens {
			claims, err := services.ValidateToken(tokenString)
			if err != nil {
				continue
			}

			c.Set("userID", claims.UserID)
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"detail": "invalid or expired token",
		})
	}
}
