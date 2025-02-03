package Admin

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
)

func ProductViewAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
		var products models.Product
		defer cancel()

		userRole := c.GetString("role")
		if userRole != models.ROLE_ADMIN && userRole != models.ROLE_SELLER {
			c.JSON(http.StatusForbidden, gin.H{"Error": "Permission denied"})
			return
		}

		storeID := c.Param("store_id")
		if storeID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
			return
		}

		objStoreID, err := primitive.ObjectIDFromHex(storeID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID format"})
			return
		}

		if userRole == models.ROLE_ADMIN {
			userID := c.GetString("user_id")
			var store models.Store
			err := global.StoreCollection.FindOne(ctx, bson.M{
				"store_id": objStoreID,
				"owner_id": userID,
			}).Decode(&store)

			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to add products to this store"})
				return
			}
		}

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

		var store models.Store
		err = global.StoreCollection.FindOne(ctx, bson.M{"name": products.Store_Name}).Decode((&store))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Không tìm thấy của hàng"})
			return
		}

		products.Store_ID = store.Store_Id
		products.Product_ID = primitive.NewObjectID()

		_, anyerr := global.ProductCollection.InsertOne(ctx, products)
		if anyerr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Created"})
			return
		}

		c.JSON(http.StatusOK, "Successfully added our Product Admin!!")
	}
}
