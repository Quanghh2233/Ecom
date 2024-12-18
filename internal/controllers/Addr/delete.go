package Addr

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("userid")
		address_id := c.Query("addressid")

		if user_id == "" || address_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusBadRequest, gin.H{"Error": "UserID or AdressID is missing"})
			c.Abort()
			return
		}
		userObjID, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid UserID"})
			return
		}

		addressObjID, err := primitive.ObjectIDFromHex(address_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid AddressID"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		filter := bson.M{"_id": userObjID}
		update := bson.M{"$pull": bson.M{"address": bson.M{"_id": addressObjID}}}

		result, err := UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to delete address"})
			return
		}

		if result.ModifiedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Address not found or already deleted"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Message": "Successfully Deleted"})
	}
}
