package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *ProductHandler) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			log.Printf("User role not found in context")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User role not found"})
			return
		}
		log.Printf("User role: %v", userRole)

		var newProduct models.Product
		if err := c.BindJSON(&newProduct); err != nil {
			log.Printf("Failed to bind JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		productID, err := h.productService.CreateProduct(ctx, userRole.(string), &newProduct)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Product added successfully",
			"product": newProduct,
			"id":      productID,
		})
	}
}
