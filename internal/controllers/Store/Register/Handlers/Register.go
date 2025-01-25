package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *StoreHandler) RegisterSeller() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("user_id")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "User ID is required",
			})
			return
		}

		var store models.Store
		if err := c.ShouldBindJSON(&store); err != nil {
			log.Printf("Failed to parse request body: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Invalid input data: " + err.Error(),
			})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		storeID, err := h.storeService.RegisterStore(ctx, userID, &store)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": "Store registered successfully",
			"storeId": storeID,
			"store":   store,
		})
	}
}
