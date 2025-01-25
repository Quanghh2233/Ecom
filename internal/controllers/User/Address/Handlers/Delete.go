package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *AddressHandler) DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("userid")
		addressID := c.Query("addressid")

		if userID == "" || addressID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusBadRequest, gin.H{"Error": "UserID or AddressID is missing"})
			c.Abort()
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := h.addressService.DeleteAddress(ctx, userID, addressID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Message": "Successfully Deleted"})
	}
}
