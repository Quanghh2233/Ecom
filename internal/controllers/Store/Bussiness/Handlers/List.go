package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) ListProductsByStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		storeID := c.Param("store_id")
		products, err := h.productService.ListProductsByStore(ctx, storeID)
		if err != nil {
			log.Printf("Error finding products: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"store_id": storeID,
			"products": products,
			"count":    len(products),
		})
	}
}
