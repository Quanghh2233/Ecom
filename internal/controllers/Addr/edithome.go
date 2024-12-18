package Addr

import (
	"context"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "User ID is missing"})
			c.Abort()
			return
		}
		userObjID, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
			return
		}

		var editAddress models.Address
		if err := c.BindJSON(&editAddress); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		filter := bson.M{"_id": userObjID, "address.type": "home"}
		update := bson.M{"$set": bson.M{
			"address.$.house_name": editAddress.House,
			"address.$.street":     editAddress.Street,
			"address.$.city_name":  editAddress.City,
			"address.$.pin_code":   editAddress.Pincode,
		}}

		result, err := UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update address"})
			return
		}

		if result.ModifiedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Home address not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Message": "Successfully updated"})
	}
}
