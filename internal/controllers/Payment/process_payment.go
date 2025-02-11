package payment

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
)

type PaymentRequest struct {
	UserID        string  `json:"user_id" binding:"required"`
	TotalAmount   float64 `json:"total_amount" binding:"required,gt=0"`
	PaymentMethod string  `json:"payment_method" binding:"required"`
}

type PaymentService interface {
	ProcessPayment(PaymentRequest) error
	UpdateInventory([]models.ProdutUser) error
}

type defaultPaymentService struct{}

func (s *defaultPaymentService) ProcessPayment(req PaymentRequest) error {
	// Implement actual payment processing here
	return nil
}

func (s *defaultPaymentService) UpdateInventory(cart []models.ProdutUser) error {
	// Implement inventory update logic here
	return nil
}

func NewPaymentService() PaymentService {
	return &defaultPaymentService{}
}

func ProcessPayment(paymentService PaymentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req PaymentRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("Invalid payment request: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment request"})
			return
		}

		if err := validatePaymentRequest(req); err != nil {
			log.Printf("Payment validation failed: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := paymentService.ProcessPayment(req); err != nil {
			log.Printf("Payment processing failed: %v", err)
			c.JSON(http.StatusPaymentRequired, gin.H{"error": "Payment processing failed"})
			return
		}

		userCart := []models.ProdutUser{} // This should be populated with actual cart items
		if err := paymentService.UpdateInventory(userCart); err != nil {
			log.Printf("Inventory update failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Inventory update failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully"})
	}
}

func validatePaymentRequest(req PaymentRequest) error {
	if req.UserID == "" {
		return fmt.Errorf("user ID is required")
	}
	if req.TotalAmount <= 0 {
		return fmt.Errorf("invalid payment amount")
	}
	if req.PaymentMethod == "" {
		return fmt.Errorf("payment method is required")
	}
	return nil
}
