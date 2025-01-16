package helper

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/database"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")

func SeedAdminUser() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Kiểm tra xem đã có admin chưa

	count, err := UserCollection.CountDocuments(ctx, bson.M{"role.name": models.ROLE_ADMIN})
	if err != nil {
		log.Fatal(err)
	}

	// Nếu chưa có admin nào, tạo admin đầu tiên
	if count == 0 {
		adminEmail := os.Getenv("ADMIN_EMAIL")       // Đọc từ env
		adminPassword := os.Getenv("ADMIN_PASSWORD") // Đọc từ env

		if adminEmail == "" || adminPassword == "" {
			log.Fatal("ADMIN_EMAIL hoặc ADMIN_PASSWORD không được thiết lập")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Không thể mã hóa mật khẩu:", err)
		}

		role, _ := NewRole(models.ROLE_ADMIN, "System Administrator")

		admin := models.User{
			First_Name: &[]string{"Admin"}[0],
			LastName:   &[]string{"System"}[0],
			Email:      &adminEmail,
			Password:   &[]string{string(hashedPassword)}[0],
			Role:       role,
			Create_At:  time.Now(),
			Update_At:  time.Now(),
			// Set các field cần thiết khác
		}

		_, err = UserCollection.InsertOne(ctx, admin)
		if err != nil {
			log.Fatal("Failed to seed admin user:", err)
		} else {
			log.Println("Admin user seeded successfully.")
		}
	}
}
