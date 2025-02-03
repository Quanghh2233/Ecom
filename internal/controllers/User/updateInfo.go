package user

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type UpdateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

func (app *Application) UpdateUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("uid")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user id not found"})
			return
		}

		var req UpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		update := bson.M{
			"$set": bson.M{
				"first_name":      req.FirstName,
				"last_name":       req.LastName,
				"phone":           req.Phone,
				"address_details": req.Address,
				"updated_at":      time.Now(),
			},
		}

		result, err := app.userCollection.UpdateOne(
			ctx,
			bson.M{"user_id": userID},
			update,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
			return
		}

		if result.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "user information updated successfully",
		})
	}
}
