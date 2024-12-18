package Addr

import (
	"context"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/database"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid code"})
			c.Abort()
			return
		}
		userObjID, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var newAddress models.Address
		newAddress.Address_id = primitive.NewObjectID()

		if err := c.BindJSON(&newAddress); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address format"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		matchFilter := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: userObjID}}}}
		projectStage := bson.D{{Key: "$project", Value: bson.D{{Key: "addressCount", Value: bson.D{{Key: "$size", Value: "$address"}}}}}}

		cursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{matchFilter, projectStage})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve address count"})
			return
		}

		var result []bson.M
		if err = cursor.All(ctx, &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process address count"})
			return
		}

		var addressCount int32
		if len(result) > 0 {
			addressCount = result[0]["addressCount"].(int32)
		}

		if addressCount >= 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not Allow "})
			return
		}

		filter := bson.M{"_id": userObjID}
		update := bson.M{"$push": bson.M{"address": newAddress}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to add address"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Address added successfully"})
	}
}
