package Auth

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	generate "github.com/Quanghh2233/Ecommerce/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

var Validate = validator.New()

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			log.Printf("[Login] Error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var founduser models.User
		err := global.UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		if err != nil {
			log.Printf("[Login] User not found: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
			return
		}

		if founduser.Password == nil || user.Password == nil {
			log.Printf("[Login] Password is missing for user: %v", user.Email)
			c.JSON(http.StatusBadRequest, gin.H{"error": "password is missing"})
			return
		}

		PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
		if !PasswordIsValid {
			log.Printf("[Login] Invalid password for user: %v, Error: %s", user.Email, msg)
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}

		if founduser.Role == nil {
			log.Printf("[Login] User role is missing for user: %v", user.Email)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user role is missing"})
			return
		}

		log.Printf("User role from database: %s", founduser.Role.Name)

		if founduser.Role.Name == "ADMIN" && founduser.User_ID == "" {
			founduser.User_ID = "admin_001" // Set a default admin ID
			log.Printf("[Login] Setting default admin ID: %s", founduser.User_ID)
		}

		token, refreshtoken, err := generate.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.LastName, founduser.User_ID, founduser.Role.Name)
		if err != nil {
			log.Printf("[Login] Error generating tokens: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = generate.UpdateAllToken(token, refreshtoken, founduser.User_ID)
		if err != nil {
			log.Printf("[Login] ID: %v", founduser.User_ID)
			log.Printf("[Login] Error updating tokens in database: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update tokens"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": founduser})
	}
}
