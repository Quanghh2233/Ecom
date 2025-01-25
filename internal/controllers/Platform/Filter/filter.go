package Adm

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		sortPrice := c.Query("sort_price") // asc or desc

		filter := bson.M{}
		opts := options.Find()

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

		// Call the sort function
		applySortPrice(opts, sortPrice)

		cursor, err := global.ProductCollection.Find(ctx, filter, opts)
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

// applySortPrice applies sorting by price to the MongoDB find options
func applySortPrice(opts *options.FindOptions, sortPrice string) {
	if sortPrice != "" {
		sortDirection := 1 // default ascending
		if sortPrice == "desc" {
			sortDirection = -1
		}
		opts.SetSort(bson.D{{Key: "price", Value: sortDirection}})
	}
}
