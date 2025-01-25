package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *OrderHandler) CancelOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("userid")
		if userQueryID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "mã người dùng trống"})
			return
		}

		orderQueryID := c.Query("orderid")
		if orderQueryID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "mã đơn hàng trống"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := h.orderService.CancelOrder(ctx, userQueryID, orderQueryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Đơn hàng đã được hủy thành công"})
	}
}

func (h *OrderHandler) CancelAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("userid")
		if userQueryID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "mã người dùng trống"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := h.orderService.CancelAllOrders(ctx, userQueryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Tất cả đơn hàng đã được hủy thành công"})
	}
}
