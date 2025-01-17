package search

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func SearchStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		var searchStores []models.Store

		queryParam := c.Query("name")
		if queryParam == "" {
			log.Println("Query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search query"})
			c.Abort()
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		search, err := global.StoreCollection.Find(ctx, bson.M{
			"$or": []bson.M{
				{"name": bson.M{"$regex": queryParam, "$options": "i"}},
				{"description": bson.M{"$regex": queryParam, "$options": "i"}},
			},
		})

		if err != nil {
			log.Printf("Error while querying MongoDB: %v", err)
			c.IndentedJSON(http.StatusNotFound, gin.H{"Error": "Something went wrong while fetching the data"})
			return
		}

		err = search.All(ctx, &searchStores)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusNotFound, gin.H{"Error": "Failed to decode search results"})
			return
		}

		defer search.Close(ctx)

		if err := search.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Invalid request"})
			return
		}

		c.IndentedJSON(http.StatusOK, searchStores)
		// log.Printf("Database: %s", Store.StoreCollection.Database().Name())
		// log.Printf("Collection: %s", Store.StoreCollection.Name())
	}
}
