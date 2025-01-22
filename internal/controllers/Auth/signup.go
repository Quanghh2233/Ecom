package Auth

import (
	"context"
	"net/http"
	"time"

	helper "github.com/Quanghh2233/Ecommerce/internal/Helper"
	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	generate "github.com/Quanghh2233/Ecommerce/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check for existing email
		count, err := global.UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking email"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}

		// Check for existing phone
		if user.Phone != nil {
			count, err = global.UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking phone"})
				return
			}
			if count > 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number already in use"})
				return
			}
		}

		// Create role
		role, err := helper.NewRole(models.DEFAULT_ROLE, "Buyer account")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user role"})
			return
		}

		// Hash password
		if user.Password != nil {
			hashedPassword := HashPassword(*user.Password)
			user.Password = &hashedPassword
		}

		// Set up user metadata
		now := time.Now()
		user.Create_At = now
		user.Update_At = now
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		user.Role = role

		// Generate tokens
		if user.Email != nil && user.First_Name != nil && user.LastName != nil {
			token, refreshToken, err := generate.TokenGenerator(
				*user.Email,
				user.Role.Name,
				*user.First_Name,
				*user.LastName,
				user.User_ID,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating tokens"})
				return
			}
			user.Token = &token
			user.Refresh_Token = &refreshToken
		}

		// Initialize empty slices
		user.UserCart = make([]models.ProdutUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)

		// Insert user into database
		_, err = global.UserCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Successfully signed up!"})
	}
}
