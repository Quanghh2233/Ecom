// filepath: /home/qhh/advance/project/Ecom/internal/controllers/Store/Addproduct.go
package Store

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		// Retrieve and log user role from context
		userRole, exists := c.Get("role")
		if !exists {
			log.Printf("User role not found in context")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User role not found"})
			return
		}
		log.Printf("User role: %v", userRole)

		// Check if the user has the required role
		if userRole != models.ROLE_ADMIN && userRole != models.ROLE_SELLER {
			log.Printf("User does not have the required role to create a product")
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have the required role to create a product"})
			return
		}

		// Bind JSON request body
		var newProduct models.Product
		if err := c.BindJSON(&newProduct); err != nil {
			log.Printf("Failed to bind JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Validate required fields
		if err := validateProduct(newProduct); err != nil {
			log.Printf("Validation error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if store exists using store_id
		var store models.Store
		err := global.StoreCollection.FindOne(ctx, bson.M{"store_id": newProduct.Store_ID}).Decode(&store)
		if err != nil {
			log.Printf("Store not found: %v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
			return
		}

		// Set product ID and store name
		newProduct.Product_ID = primitive.NewObjectID()
		newProduct.Store_Name = store.Name // Ensure store name matches the found store

		// Insert product
		result, err := global.ProductCollection.InsertOne(ctx, newProduct)
		if err != nil {
			log.Printf("Failed to create product: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Product created successfully",
			"product": newProduct,
			"id":      result.InsertedID,
		})
	}
}

// Helper function to validate product
func validateProduct(product models.Product) error {
	if product.Product_Name == "" {
		return fmt.Errorf("product name is required")
	}
	if product.Price == nil || *product.Price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}
	if product.Quantity < 0 {
		return fmt.Errorf("quantity cannot be negative")
	}
	if product.Store_ID.IsZero() {
		return fmt.Errorf("store_id is required")
	}
	return nil
}
