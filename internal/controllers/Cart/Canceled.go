package Cart

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *Application) CancelList() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("uid")

		if userID == "" {
			log.Println("Error: User ID not found")
			c.JSON(http.StatusBadRequest, gin.H{"error": "user ID not found"})
			return
		}

		objID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			log.Printf("Error: Invalid user ID format: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		filter := bson.M{
			"user_id": objID,
			"status":  "CANCELED",
		}

		cursor, err := app.orderCollection.Find(ctx, filter)
		if err != nil {
			log.Printf("Error: failed to fetch canceled orders for user %s: %v", userID, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch canceled orders"})
			return
		}
		defer cursor.Close(ctx)

		var canceledOrders []models.Order
		if err := cursor.All(ctx, &canceledOrders); err != nil {
			log.Printf("Error: Failed to decode cancel orders: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode canceled orders"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"orders": canceledOrders,
		})
	}
}
