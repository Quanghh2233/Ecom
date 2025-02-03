package order

import (
	"context"
	"net/http"
	"time"

	cart "github.com/Quanghh2233/Ecommerce/internal/database/Cart"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("pid")
		if productQueryID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product id is empty"})
			return
		}

		userQueryID := c.Query("userid")
		if userQueryID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user id is empty"})
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id format"})
			return
		}

		userID, err := primitive.ObjectIDFromHex(userQueryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id format"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = cart.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userID.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successfully placed the order"})
	}
}
