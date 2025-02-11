package user

import (
	"context"
	"log"
	"net/http"
	"time"

	helper "github.com/Quanghh2233/Ecommerce/internal/Helper"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *Application) ChangeRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("uid")

		if userID == "" {
			log.Println("Error: user id not found in context")
			c.JSON(http.StatusBadRequest, gin.H{"error": "user id not found"})
			return
		}

		// Kiểm tra xem userCollection có bị nil không
		if app.userCollection == nil {
			log.Println("Error: userCollection is nil")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database collection is not initialized"})
			return
		}
		// Chuyển userID từ string sang ObjectID
		objID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			log.Printf("Error: Invalid userID format: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
			return
		}

		// Tạo context với timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Kiểm tra xem user có tồn tại không
		var user models.User
		err = app.userCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
		if err != nil {
			log.Printf("Error: User not found with ID %s: %v", userID, err)
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		var newRole *models.Role
		var roleName string

		// //Điều kiện
		hasStore, err := app.CheckUserStore(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check store ownership"})
			return
		}

		if !hasStore {
			log.Println("Error: User does not own any store, cannot become SELLER")
			c.JSON(http.StatusBadRequest, gin.H{"error": "User must own at least one store to become SELLER"})
			return
		}

		// Toggle giữa BUYER <-> SELLER
		if user.Role.Name == models.ROLE_SELLER {
			roleName = "BUYER"
			newRole, err = helper.NewRole(models.ROLE_BUYER, roleName)
			log.Println("Debug: Changing role to BUYER")
		} else {
			roleName = "SELLER"
			newRole, err = helper.NewRole(models.ROLE_SELLER, roleName)
			log.Println("Debug: Changing role to SELLER")
		}

		if err != nil {
			log.Printf("Error: Failed to create new role: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
			return
		}

		// Cập nhật role trong database
		update := bson.M{"$set": bson.M{"role": newRole}}
		updateResult, err := app.userCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
		if err != nil {
			log.Printf("Error: Failed to update role for user %s: %v", userID, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role"})
			return
		}

		// err = token.InvalidateSession(userID, app.redisClient)
		// if err != nil {
		// 	log.Printf("Error: Failed to invalidate session: %v", err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invalidate session"})
		// 	return
		// }

		// log.Println("Debug: User session invalidated, user must re-login")
		if updateResult.ModifiedCount == 0 {
			log.Printf("Warning: No document updated for user %s", userID)
		}

		// Trả về phản hồi JSON
		c.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"message":  "User role updated",
			"new_role": newRole.Name,
		})
	}
}
