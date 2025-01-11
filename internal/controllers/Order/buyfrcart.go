package order

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	cart "github.com/Quanghh2233/Ecommerce/internal/database/Cart"
	"github.com/gin-gonic/gin"
)

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")

		if userQueryID == "" {
			log.Panicln("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("UserID is empty"))
		}

		var selectedItems []string
		if err := c.ShouldBindJSON(&selectedItems); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
			return
		}

		if len(selectedItems) == 0 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No item selected"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		orderList, err := cart.BuyItemFromCart(ctx, app.userCollection, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error":  "Failed to place order",
				"detail": err.Error(),
			})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"message":    "Successfully placed the order",
			"order_list": orderList,
		})
	}
}
