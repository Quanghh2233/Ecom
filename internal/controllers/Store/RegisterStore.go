package Store

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *Application) RegisterSeller() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy thông tin userID từ query
		userQueryID := c.Query("user_id")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		// Parse dữ liệu từ request body
		var shop models.Store
		if err := c.ShouldBindJSON(&shop); err != nil {
			log.Println("Invalid input:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
			return
		}

		if shop.Name == "" || shop.Email == "" || shop.Phone == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}

		// Gắn thông tin chủ shop (người dùng)
		// firstName := c.GetString("first_name")
		// lastName := c.GetString("last_name")
		// ownerName := firstName + " " + lastName
		// if ownerName == " " { // Nếu không có thông tin, đặt giá trị mặc định
		// 	ownerName = "Anonymous"
		// }
		// shop.Owner = ownerName
		shop.Owner = userQueryID
		shop.Store_Id = primitive.NewObjectID()
		shop.CreateAt = time.Now()

		// Kết nối database và lưu thông tin shop
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := app.storeCollection.InsertOne(ctx, shop)
		if err != nil {
			log.Println("Failed to register shop:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register shop"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Shop registered successfully",
			"shop":    shop,
		})
	}
}
