package Cart

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	cart "github.com/Quanghh2233/Ecommerce/internal/database/Cart"
	"github.com/gin-gonic/gin"
)

func (app *Application) GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("userid")
		if userQueryID == "" {
			log.Println("[Orders] Error: User ID is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		log.Printf("[Orders] Operation: GetOrders | User: %s | Status: Starting", userQueryID)

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		orders, err := cart.GetUserOrders(ctx, app.userCollection, userQueryID)
		if err != nil {
			log.Printf("[Orders] Operation: GetOrders | User: %s | Status: Failed | Error: %v", userQueryID, err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.Printf("[Orders] Operation: GetOrders | User: %s | Status: Success | Orders Count: %d", userQueryID, len(orders))
		c.IndentedJSON(http.StatusOK, gin.H{
			"user_id": userQueryID,
			"orders":  orders,
		})
	}
}
