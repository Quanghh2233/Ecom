package Store

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		// Get store ID and convert to ObjectID
		storeID := c.Param("store_id")
		objID, err := primitive.ObjectIDFromHex(storeID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
			return
		}

		// Get store details
		var store models.Store
		err = global.StoreCollection.FindOne(ctx, bson.M{"store_id": objID}).Decode(&store)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve store"})
			}
			return
		}

		// Get pagination parameters
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		skip := (page - 1) * limit

		// Setup options for pagination and sorting
		findOptions := options.Find().
			SetSkip(int64(skip)).
			SetLimit(int64(limit)).
			SetSort(bson.D{{Key: "create_at", Value: 1}})

		// Query products
		cursor, err := global.ProductCollection.Find(ctx,
			bson.M{"store_id": objID},
			findOptions,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}
		defer cursor.Close(ctx)

		var products []bson.M
		if err := cursor.All(ctx, &products); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode products"})
			return
		}

		// Get total products count
		total, err := global.ProductCollection.CountDocuments(ctx, bson.M{"store_id": objID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count products"})
			return
		}

		response := StoreResponse{
			Store:    store,
			Products: products,
			Message:  "Store retrieved successfully",
		}
		response.Pagination.Total = total
		response.Pagination.Page = page
		response.Pagination.Limit = limit
		response.Pagination.Pages = (total + int64(limit) - 1) / int64(limit)

		c.JSON(http.StatusOK, response)
	}
}

type StoreResponse struct {
	Store      models.Store `json:"store"`
	Products   []bson.M     `json:"products"`
	Pagination struct {
		Total int64 `json:"total"`
		Page  int   `json:"page"`
		Limit int   `json:"limit"`
		Pages int64 `json:"pages"`
	} `json:"pagination"`
	Message string `json:"message"`
}
