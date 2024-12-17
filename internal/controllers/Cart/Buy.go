package Cart

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryID := c.Query("id")

		if userQueryID == "" {
			log.Panicln("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("UserID is empty"))
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		err := database.BuyItemFromCart(ctx, app.userCollection, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully placed the order")
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("pid")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))

			return
		}

		userQueryID := c.Query("userid")
		if userQueryID == "" {
			log.Println("user id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)

		}

		c.IndentedJSON(200, "Successfully placed the order")
	}
}

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
