package Adm

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/database"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProductViewAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
		var products models.Product
		defer cancel()
		if err := c.BindJSON(&products); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if products.Store_Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "store_name là bắt buộc"})
			return
		}

		// if products.Store_ID.IsZero() {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		// 	return
		// }

		storeCollection := database.Client.Database("Ecommerce").Collection("Stores")
		var store models.Store
		err := storeCollection.FindOne(ctx, bson.M{"name": products.Store_Name}).Decode((&store))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Không tìm thấy của hàng"})
			return
		}

		products.Store_ID = store.Store_Id
		products.Product_ID = primitive.NewObjectID()

		_, anyerr := ProductCollection.InsertOne(ctx, products)
		if anyerr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Created"})
			return
		}

		c.JSON(http.StatusOK, "Successfully added our Product Admin!!")
	}
}
