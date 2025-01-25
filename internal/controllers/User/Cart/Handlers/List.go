package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	service "github.com/Quanghh2233/Ecommerce/internal/controllers/User/Cart/Service"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) GetOrders() gin.HandlerFunc {
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

		orders, err := h.orderService.GetUserOrders(ctx, userQueryID)
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
