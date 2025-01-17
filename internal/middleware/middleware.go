package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/Quanghh2233/Ecommerce/internal/token"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")
		if ClientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization Header Provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(ClientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Set("role", claims.Role)
		log.Printf("Authenticated User Role: %s", claims.Role)

		c.Next()
	}
}

func AuthRole(requireRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		userRole := strings.ToLower(role.(string))
		for _, r := range requireRole {
			if strings.ToLower(r) == userRole {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}
