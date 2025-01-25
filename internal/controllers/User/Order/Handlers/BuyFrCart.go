package handlers

import (
	"context"
	"net/http"
	"time"

	service "github.com/Quanghh2233/Ecommerce/internal/controllers/User/Order/Service"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService service.CartService
}

func NewCartHandler(cartService service.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}

func (h *CartHandler) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")
		if userQueryID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "mã người dùng trống"})
			return
		}

		var selectedItems []string
		if err := c.ShouldBindJSON(&selectedItems); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Định dạng đầu vào không hợp lệ"})
			return
		}

		if len(selectedItems) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Không có sản phẩm nào được chọn"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderList, err := h.cartService.BuyItemsFromCart(ctx, userQueryID, selectedItems)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":  "Không thể đặt hàng",
				"detail": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "Đặt hàng thành công",
			"order_list": orderList,
		})
	}
}
