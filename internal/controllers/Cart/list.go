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
		// Lấy user ID từ query
		userQueryID := c.Query("userid")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Gọi hàm từ database để lấy danh sách đơn hàng
		orders, err := cart.GetUserOrders(ctx, app.userCollection, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Trả về danh sách đơn hàng
		c.IndentedJSON(http.StatusOK, gin.H{
			"user_id": userQueryID,
			"orders":  orders,
		})
	}
}
