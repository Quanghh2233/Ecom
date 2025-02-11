package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	token "github.com/Quanghh2233/Ecommerce/internal/token"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		refreshToken := c.Request.Header.Get("refresh-token")

		if clientToken == "" && refreshToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization Token Provided"})
			c.Abort()
			return
		}

		var claims *token.SignedDetails
		var errMessage string

		// ctx := context.Background()
		// _, redisErr := database.RedisClient.Get(ctx, "token:"+clientToken).Result()

		// if redisErr == nil {
		// 	claims, errMessage = token.ValidateToken(clientToken)
		// } else {
		// 	claims, errMessage = token.ValidateToken(refreshToken)
		// 	if errMessage == "" { // Đổi từ err != "" sang err == "" vì đang kiểm tra không có lỗi
		// 		_, newToken, genErr := token.TokenGenerator(claims.Uid, claims.First_name, claims.Last_name, claims.Email, claims.Role)
		// 		if genErr != nil {
		// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo token"})
		// 			c.Abort()
		// 			return
		// 		}

		// 		// Không khai báo lại err, sử dụng err đã có
		// 		setErr := database.RedisClient.Set(ctx, "token:"+newToken, "active", 15*time.Minute).Err()
		// 		if setErr != nil {
		// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu token vào Redis"})
		// 			c.Abort()
		// 			return
		// 		}

		// 		c.Header("New-token", newToken)
		// 	}
		// }
		if clientToken != "" {
			claims, errMessage = token.ValidateToken(clientToken)
		} else {
			claims, errMessage = token.ValidateToken(refreshToken)
		}

		if errMessage != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errMessage})
			c.Abort()
			return
		}

		// Skip further checks for admin users
		if claims.Role == models.ROLE_ADMIN {
			c.Set("email", claims.Email)
			c.Set("uid", claims.Uid)
			c.Set("role", claims.Role)
			c.Next()
			return
		}

		if (claims.Email == "" || claims.Uid == "") && claims.Role == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: missing required claims"})
			c.Abort()
			return
		}

		user, dberr := getUserFromDatabase(claims.Uid)
		if dberr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Set("role", claims.Role)
		c.Set("user", user)
		c.Next()
	}
}

func getUserFromDatabase(ID string) (models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := global.UserCollection.FindOne(ctx, bson.M{"user_id": ID}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CheckPermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			log.Println("User:", role)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		// Skip permission check for admin
		if role == models.ROLE_ADMIN {
			c.Next()
			return
		}

		user, exists := c.Get("user")
		if !exists {
			log.Println("User:", user)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		authUser, ok := user.(models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			c.Abort()
			return
		}

		if authUser.Role == nil || !authUser.Role.HasPermission(permission) {
			log.Printf("role: %+v", authUser.Role)
			log.Println(authUser.Role.HasPermission(permission))
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
