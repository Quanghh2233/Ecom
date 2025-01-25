package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *OrderHandler) InstantBuy() gin.HandlerFunc {
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

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := h.orderService.InstantBuy(ctx, productQueryID, userQueryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successfully placed the order"})
	}
}
