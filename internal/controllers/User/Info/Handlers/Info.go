package handlers

import (
	"context"
	"net/http"
	"time"

	service "github.com/Quanghh2233/Ecommerce/internal/controllers/User/User/Service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("uid")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user id not found"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		profile, err := h.userService.GetUserProfile(ctx, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching user profile"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"user":   profile,
		})
	}
}
