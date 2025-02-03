package Cart

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("user_id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			log.Println("user_id:", userID)
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid id"})
			c.Abort()
			return
		}

		userObjectID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			log.Printf("Invalid user ID format: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User
		err = global.UserCollection.FindOne(ctx, bson.M{"_id": userObjectID}).Decode(&user)
		if err != nil {
			log.Printf("Error finding user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			return
		}

		filterMatch := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: userObjectID}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$usercart"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{{Key: "$sum", Value: "$usercart.price"}}}}}}

		cursor, err := global.UserCollection.Aggregate(ctx, mongo.Pipeline{filterMatch, unwind, group})
		if err != nil {
			log.Printf("Error aggregating cart items: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to aggregate cart items"})
			return
		}

		var results []bson.M
		if err = cursor.All(ctx, &results); err != nil {
			log.Printf("Error decoding aggregation results: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode aggregation results"})
			return
		}

		if len(results) == 0 {
			c.JSON(http.StatusOK, gin.H{"total": 0, "cart": user.UserCart})
			return
		}

		total := results[0]["total"]
		c.JSON(http.StatusOK, gin.H{"total": total, "cart": user.UserCart})
	}
}
