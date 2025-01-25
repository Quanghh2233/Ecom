package Store

import (
	"context"
	"log"
	"net/http"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListProductsByStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		// Get store ID from URL parameter
		storeID := c.Param("store_id")
		objID, err := primitive.ObjectIDFromHex(storeID)
		if err != nil {
			log.Printf("Invalid store ID: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
			return
		}

		// Query products by store_id
		cursor, err := global.ProductCollection.Find(ctx, bson.M{"store_id": objID})
		if err != nil {
			log.Printf("Error finding products: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}
		defer cursor.Close(ctx)

		var products []bson.M
		if err := cursor.All(ctx, &products); err != nil {
			log.Printf("Error decoding products: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode products"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"store_id": storeID,
			"products": products,
			"count":    len(products),
		})
	}
}
