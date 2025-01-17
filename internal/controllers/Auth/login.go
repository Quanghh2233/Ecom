package Auth

import (
	"context"
	"fmt"
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
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		var founduser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := global.UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
			return
		}

		if founduser.Password == nil || user.Password == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "password is missing"})
			return
		}

		PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
		defer cancel()

		if !PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}

		if founduser.Role == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user role is missing"})
			return
		}

		log.Printf("User role from database: %s", founduser.Role.Name)

		token, refreshtoken, _ := generate.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.LastName, founduser.User_ID, founduser.Role.Name)
		defer cancel()

		generate.UpdateAllToken(token, refreshtoken, founduser.User_ID)

		c.JSON(http.StatusFound, founduser)
	}
}
