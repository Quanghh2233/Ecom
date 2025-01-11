package Ecom

import (
	"context"
	"net/http"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newProduct models.Product
		if err := c.BindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		var store models.Store
		err := StoreCollection.FindOne(context.Background(), bson.M{"store_id": newProduct.Store_ID}).Decode(&store)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
			return
		}

		newProduct.Product_ID = primitive.NewObjectID()

		_, err = ProductCollection.InsertOne(context.Background(), newProduct)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Product created successfully",
			"product": newProduct,
		})
	}
}
