package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	service "github.com/Quanghh2233/Ecommerce/internal/controllers/Store/Bussiness/Service"
	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		userRole := c.GetString("role")
		userID := c.GetString("user_id")
		storeID := c.Param("store_id")
		productID := c.Param("product_id")

		var updateData service.UpdateProductData
		if err := c.BindJSON(&updateData); err != nil {
			log.Printf("[UpdateProduct] Invalid request body for User ID: %s, Error: %v", userID, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
			return
		}

		err := h.productService.UpdateProduct(ctx, userRole, userID, storeID, productID, &updateData)
		if err != nil {
			log.Printf("[UpdateProduct] Failed to update Product ID: %s, Error: %v", productID, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":        "Product updated successfully",
			"updated_fields": updateData,
		})
	}
}
