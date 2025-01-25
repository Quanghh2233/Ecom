package Adm

import (
	"context"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DelMultiple() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var product_id []string
		if err := c.ShouldBindJSON(&product_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
			return
		}

		if len(product_id) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
			return
		}

		var objIDs []primitive.ObjectID
		for _, id := range product_id {
			objID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID:" + id})
				return
			}
			objIDs = append(objIDs, objID)
		}

		filter := bson.M{"product_id": bson.M{"$in": objIDs}}
		result, err := global.ProductCollection.DeleteMany(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete products"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":      "products deleted successfully",
			"delete_count": result.DeletedCount,
		})
	}
}
