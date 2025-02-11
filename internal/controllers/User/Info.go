package user

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (app *Application) GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userID := c.GetString("uid")
		if userID == "" {
			userID = c.Query("user_id")
		}

		if userID == "" {
			log.Println("[GetUserInfo] User ID missing")
			c.JSON(http.StatusBadRequest, gin.H{"error": "user ID is required"})
			return
		}

		// Lấy thông tin user từ MongoDB
		var user models.User
		err := global.UserCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				log.Printf("[GetUserInfo] User not found: %v", userID)
				c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			} else {
				log.Printf("[GetUserInfo] Database error: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			}
			return
		}

		// Truy vấn danh sách store của user
		cursor, err := app.storeCollection.Find(ctx, bson.M{"owner": userID})
		if err != nil {
			log.Printf("[GetUserInfo] Error fetching stores: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching stores"})
			return
		}

		var stores []models.Store
		if err = cursor.All(ctx, &stores); err != nil {
			log.Printf("[GetUserInfo] Error decoding stores: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error decoding stores"})
			return
		}

		log.Printf("[GetUserInfo] Found %d stores for user %s", len(stores), userID)

		// Kiểm tra xem user có store không
		hasStores, err := app.CheckUserStore(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check store ownership"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"user": gin.H{
				"name":     *user.First_Name + " " + *user.LastName,
				"email":    user.Email,
				"phone":    user.Phone,
				"address":  user.Address_Details,
				"cart":     user.UserCart,
				"orders":   user.Order_Status,
				"stores":   stores,
				"hasStore": hasStores,
			},
		})
	}
}
