package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *CartHandler) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy user ID từ token claims thay vì query
		userID, exists := c.Get("uid")
		if !exists {
			log.Printf("[Cart] Error: User not authenticated")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		productQueryID := c.Query("product_id")
		if productQueryID == "" {
			log.Printf("[Cart] Error: Empty product ID | Status: %d", http.StatusBadRequest)
			c.JSON(http.StatusBadRequest, gin.H{"error": "product id is empty"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		userIDStr := userID.(string)
		log.Printf("[Cart] Operation: AddToCart | User: %s | Product: %s | Status: Starting",
			userIDStr, productQueryID)

		err := h.cartService.AddProductToCart(ctx, productQueryID, userIDStr)
		if err != nil {
			log.Printf("[Cart] Operation: AddToCart | User: %s | Product: %s | Status: Failed | Error: %v",
				userIDStr, productQueryID, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.Printf("[Cart] Operation: AddToCart | User: %s | Product: %s | Status: Success",
			userIDStr, productQueryID)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully added to the cart"})
	}
}
