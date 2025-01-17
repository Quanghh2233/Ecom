package Adm

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func FilterProd() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		brand := c.Query("brand")
		category := c.Query("category")
		minPrice := c.Query("min_price")
		maxPrice := c.Query("max_price")
		minRating := c.Query("minRating")

		filter := bson.M{}

		if brand != "" {
			filter["brand"] = brand
		}

		if category != "" {
			filter["category"] = category
		}

		if minPrice != "" || maxPrice != "" {
			priceFilter := bson.M{}
			if minPrice != "" {
				minPriceVal, err := strconv.Atoi(minPrice)
				if err == nil {
					priceFilter["$gte"] = minPriceVal
				}
			}

			if maxPrice != "" {
				maxPriceVal, err := strconv.Atoi(maxPrice)
				if err == nil {
					priceFilter["$lte"] = maxPriceVal
				}
			}
			filter["price"] = priceFilter
		}

		if minRating != "" {
			filter["rating"] = bson.M{"$gte": minRating}
		}

		cursor, err := global.ProductCollection.Find(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "error fetching products"})
			return
		}
		defer cursor.Close(ctx)

		var products []bson.M
		if err = cursor.All(ctx, &products); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"filtered_products": products})
	}
}
