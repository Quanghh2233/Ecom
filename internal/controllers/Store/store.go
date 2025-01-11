package Ecom

import (
	"context"
	"net/http"

	"github.com/Quanghh2233/Ecommerce/internal/database"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var StoreCollection *mongo.Collection = database.ProductData(database.Client, "Stores")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")

func GetStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		storeID := c.Param("store_id")
		objID, err := primitive.ObjectIDFromHex(storeID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
			return
		}

		var store models.Store
		err = StoreCollection.FindOne(context.Background(), bson.M{"store_id": objID}).Decode(&store)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrive store"})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Store retrieved successfully",
			"store":   store,
		})
	}
}
