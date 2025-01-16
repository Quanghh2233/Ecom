package Store

import (
	"context"
	"log"
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
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if newProduct.Product_Name == "" || newProduct.Price <= 0 || newProduct.Quantity < 0 {
			log.Println("Missing required field or invalid value")
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Missing required field or invalid value"})
			return
		}

		if newProduct.Store_ID.IsZero() {
			log.Println("store_id is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
			return
		}

		var store models.Store
		err := StoreCollection.FindOne(context.Background(), bson.M{"store_name": newProduct.Store_Name}).Decode(&store)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
			return
		}

		newProduct.Product_ID = primitive.NewObjectID()

		_, err = ProductCollection.InsertOne(context.Background(), newProduct)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Product created successfully",
			"product": newProduct,
		})
	}
}
