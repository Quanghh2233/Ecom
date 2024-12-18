package order

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	cart "github.com/Quanghh2233/Ecommerce/internal/database/Cart"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *Application) CancelOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("userid")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		userID, err := primitive.ObjectIDFromHex(userQueryID)
		if err != nil {
			log.Println("invalid user id format:", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user id format"})
			return
		}

		orderQueryID := c.Query("orderid")
		if orderQueryID == "" {
			log.Println("order id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("order id is empty"))
			return
		}

		orderID, err := primitive.ObjectIDFromHex(orderQueryID)
		if err != nil {
			log.Println("Invalid order id format:", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid order id format"})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		filter := bson.M{
			"_id":    userID,
			"orders": bson.M{"$elemMatch": bson.M{"_id": orderID}},
		}

		update := bson.M{
			"$pull": bson.M{"orders": bson.M{"_id": orderID}},
		}

		result, err := app.userCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Println("Failed to cancel order:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel order"})
			return
		}

		if result.ModifiedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found or already cancelled"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order successfully cancelled"})
	}
}

func (app *Application) CancelAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("userid")
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}

		err := cart.CancelAllOrder(context.Background(), app.userCollection, userID)
		if err != nil {
			if err == cart.ErrUserIdIsNotValid {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel orders"})
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "All orders have been canceled"})
	}
}
