package Store

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

func UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		userRole := c.GetString("role")
		userID := c.GetString("user_id")
		log.Printf("[UpdateProduct] User Role: %s, User ID: %s", userRole, userID)

		if userRole != models.ROLE_ADMIN && userRole != models.ROLE_SELLER {
			log.Printf("[UpdateProduct] Permission denied for User ID: %s", userID)
			c.JSON(http.StatusForbidden, gin.H{"Error": "Permission denied"})
			return
		}

		storeID := c.Param("store_id")
		if storeID == "" {
			log.Printf("[UpdateProduct] Missing store_id for User ID: %s", userID)
			c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
			return
		}

		objStoreID, err := primitive.ObjectIDFromHex(storeID)
		if err != nil {
			log.Printf("[UpdateProduct] Invalid store ID format: %s", storeID)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID format"})
			return
		}

		if userRole == models.ROLE_SELLER {
			var store models.Store
			err := global.StoreCollection.FindOne(ctx, bson.M{
				"store_id": objStoreID,
				"owner_id": userID,
			}).Decode(&store)

			if err != nil {
				log.Printf("[UpdateProduct] Permission denied for User ID: %s to update Store ID: %s", userID, storeID)
				c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to add products to this store"})
				return
			}
		}

		productID := c.Param("product_id")
		if productID == "" {
			log.Printf("[UpdateProduct] Missing product_id for User ID: %s", userID)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
			return
		}

		objID, err := primitive.ObjectIDFromHex(productID)
		if err != nil {
			log.Printf("[UpdateProduct] Invalid product ID format: %s", productID)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
			return
		}

		var existingProduct models.Product
		err = global.ProductCollection.FindOne(ctx, bson.M{"product_id": objID}).Decode(&existingProduct)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				log.Printf("[UpdateProduct] Product not found for Product ID: %s", productID)
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
			log.Printf("[UpdateProduct] Error finding product: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find product", "details": err.Error()})
			return
		}

		log.Printf("[UpdateProduct] Found existing product: %+v", existingProduct)

		var updateData struct {
			ProductName *string  `json:"product_name"`
			Price       *uint64  `json:"price"`
			Rating      *float32 `json:"rating"`
			Image       *string  `json:"image"`
		}

		if err := c.BindJSON(&updateData); err != nil {
			log.Printf("[UpdateProduct] Invalid request body for User ID: %s, Error: %v", userID, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
			return
		}

		update := bson.M{"$set": bson.M{}}

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

		update["$set"].(bson.M)["updated_at"] = time.Now()

		log.Printf("[UpdateProduct] Updating Product ID: %s with Data: %+v", productID, update["$set"])

		result, err := global.ProductCollection.UpdateOne(ctx, bson.M{"product_id": objID}, update)
		if err != nil {
			log.Printf("[UpdateProduct] Failed to update Product ID: %s, Error: %v", productID, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product", "details": err.Error()})
			return
		}

		if result.MatchedCount == 0 {
			log.Printf("[UpdateProduct] Product not found for Product ID: %s", productID)
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		if result.ModifiedCount == 0 {
			log.Printf("[UpdateProduct] No changes made to Product ID: %s", productID)
			c.JSON(http.StatusOK, gin.H{"message": "No changes made to the product"})
			return
		}

		log.Printf("[UpdateProduct] Successfully updated Product ID: %s", productID)
		c.JSON(http.StatusOK, gin.H{
			"message":        "Product updated successfully",
			"updated_fields": update["$set"],
		})
	}
}
