package Store

import (
	"context"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *Application) AdmAddStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var newStore models.Store
		if err := c.BindJSON(&newStore); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newStore.Store_Id = primitive.NewObjectID()
		newStore.CreateAt = time.Now()

		result, err := app.storeCollection.InsertOne(ctx, newStore)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create store"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":  "Successfully created a new store!",
			"store_id": result.InsertedID,
			"store":    newStore,
		})
	}
}
