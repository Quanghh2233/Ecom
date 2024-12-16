package Ecom

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		productID := c.Param("product_id")
		if productID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
			return
		}

		// Convert productID to ObjectID
		objID, err := primitive.ObjectIDFromHex(productID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
			return
		}

		// Struct to receive update data
		var updateData struct {
			ProductName *string  `json:"product_name"`
			Price       *uint64  `json:"price"`
			Rating      *float32 `json:"rating"`
			Image       *string  `json:"image"`
		}

		// Bind JSON data
		if err := c.BindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
			return
		}

		// Prepare update document
		update := bson.M{"$set": bson.M{}}

		// Conditionally add fields to update
		if updateData.ProductName != nil {
			update["$set"].(bson.M)["product_name"] = updateData.ProductName
		}
		if updateData.Price != nil {
			update["$set"].(bson.M)["price"] = updateData.Price
		}
		if updateData.Rating != nil {
			update["$set"].(bson.M)["rating"] = updateData.Rating
		}
		if updateData.Image != nil {
			update["$set"].(bson.M)["image"] = updateData.Image
		}

		// Add updated timestamp
		update["$set"].(bson.M)["updated_at"] = time.Now()

		// Perform update
		result, err := ProductCollection.UpdateOne(
			ctx,
			bson.M{"_id": objID},
			update,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product", "details": err.Error()})
			return
		}

		// Check if product was found and updated
		if result.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		// Check if any fields were actually modified
		if result.ModifiedCount == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No changes made to the product"})
			return
		}

		// Success response
		c.JSON(http.StatusOK, gin.H{
			"message":        "Product updated successfully",
			"updated_fields": update["$set"],
		})
	}
}
