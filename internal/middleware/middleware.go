package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	token "github.com/Quanghh2233/Ecommerce/internal/token"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		refreshToken := c.Request.Header.Get("refresh-token")

		if clientToken == "" && refreshToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization Token Provided"})
			c.Abort()
			return
		}

		var claims *token.SignedDetails
		var err string

		if clientToken != "" {
			claims, err = token.ValidateToken(clientToken)
		} else {
			claims, err = token.ValidateToken(refreshToken)
		}

		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		// Skip further checks for admin users
		if claims.Role == models.ROLE_ADMIN {
			c.Set("email", claims.Email)
			c.Set("uid", claims.Uid)
			c.Set("role", claims.Role)
			c.Next()
			return
		}

		if (claims.Email == "" || claims.Uid == "") && claims.Role == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: missing required claims"})
			c.Abort()
			return
		}

		user, dberr := getUserFromDatabase(claims.Uid)
		if dberr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Set("role", claims.Role)
		c.Set("user", user)
		c.Next()
	}
}

func getUserFromDatabase(ID string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := global.UserCollection.FindOne(ctx, bson.M{"user_id": ID}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CheckPermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			log.Println("User:", role)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		// Skip permission check for admin
		if role == models.ROLE_ADMIN {
			c.Next()
			return
		}

		user, exists := c.Get("user")
		if !exists {
			log.Println("User:", user)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		authUser, ok := user.(models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		if authUser.Role == nil || !authUser.Role.HasPermission(permission) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
