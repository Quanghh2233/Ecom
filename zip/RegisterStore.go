package Store

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *Application) RegisterSeller() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from query
		userID := c.Query("user_id")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "User ID is required",
			})
			return
		}

		// Parse request body
		var store models.Store
		if err := c.ShouldBindJSON(&store); err != nil {
			log.Printf("Failed to parse request body: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Invalid input data: " + err.Error(),
			})
			return
		}

		// Validate required fields
		if err := validateStoreInput(&store); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		// Check for existing store
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		exists, err := app.checkStoreExists(ctx, store.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  "Database error",
			})
			return
		}
		if exists {
			c.JSON(http.StatusConflict, gin.H{
				"status": "error",
				"error":  "Store with this email already exists",
			})
			return
		}

		// Prepare store data
		store.Store_Id = primitive.NewObjectID()
		store.Owner = userID
		store.CreateAt = time.Now()
		store.Status = "active"

		// Save to database
		result, err := app.storeCollection.InsertOne(ctx, store)
		if err != nil {
			log.Printf("Failed to register store: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  "Failed to register store",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": "Store registered successfully",
			"storeId": result.InsertedID,
			"store":   store,
		})
	}
}

func validateStoreInput(store *models.Store) error {
	if store.Name == "" {
		return errors.New("store name is required")
	}
	if store.Email == "" {
		return errors.New("email is required")
	}
	if store.Phone == "" {
		return errors.New("phone number is required")
	}
	return nil
}

func (app *Application) checkStoreExists(ctx context.Context, email string) (bool, error) {
	count, err := app.storeCollection.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
