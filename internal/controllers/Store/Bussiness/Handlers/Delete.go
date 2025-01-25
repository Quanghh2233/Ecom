package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role")
		userID := c.GetString("user_id")
		storeID := c.Param("store_id")
		productID := c.Param("product_id")

		if storeID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
			return
		}

		if productID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product_id is required"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		err := h.productService.DeleteProduct(ctx, userRole, userID, storeID, productID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	}
}
