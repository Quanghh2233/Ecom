package user

import (
	"context"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (app *Application) GetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("uid")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user id not found"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User
		err := app.userCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching user"})
			return
		}

		cursor, err := app.storeCollection.Find(ctx, bson.M{"owner": userID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching stores"})
			return
		}
		defer cursor.Close(ctx)

		var stores []models.Store
		if err = cursor.All(ctx, &stores); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error decoding stores"})
			return
		}

		hasStores := app.CheckUserStore(userID)

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

func (app *Application) CheckUserStore(userID string) bool {
	if userID == "" {
		return false
	}

	if app.storeCollection == nil {
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := app.storeCollection.CountDocuments(ctx, bson.M{"owner": userID})
	if err != nil {
		return false
	}

	return count > 0
}
