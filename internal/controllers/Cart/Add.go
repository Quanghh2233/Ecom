package Cart

import (
	"context"
	"log"
	"net/http"
	"time"

	cart "github.com/Quanghh2233/Ecommerce/internal/database/Cart"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from token claims instead of query
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

		productID, err := primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Printf("[Cart] Error: Invalid product ID format | ID: %s | Error: %v", productQueryID, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		userIDStr := userID.(string)
		log.Printf("[Cart] Operation: AddToCart | User: %s | Product: %s | Status: Starting",
			userIDStr, productID)

		err = cart.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userIDStr)
		if err != nil {
			log.Printf("[Cart] Operation: AddToCart | User: %s | Product: %s | Status: Failed | Error: %v",
				userIDStr, productID, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.Printf("[Cart] Operation: AddToCart | User: %s | Product: %s | Status: Success",
			userIDStr, productID)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully added to the cart"})
	}
}
