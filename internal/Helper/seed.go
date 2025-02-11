package helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/database"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/Quanghh2233/Ecommerce/internal/token"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")

func SeedAdminUser() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Đọc thông tin admin từ environment variables
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	// Kiểm tra environment variables
	if adminEmail == "" || adminPassword == "" {
		return fmt.Errorf("ADMIN_EMAIL và ADMIN_PASSWORD phải được thiết lập trong environment variables")
	}

	// Kiểm tra xem đã có admin chưa
	var existingAdmin models.User
	err := UserCollection.FindOne(ctx, bson.M{"email": adminEmail}).Decode(&existingAdmin)
	if err == nil {
		log.Println("Admin account already exists")
		return nil
	} else if err != mongo.ErrNoDocuments {
		return fmt.Errorf("error checking existing admin: %v", err)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	// Generate tokens
	adminRole, _ := NewRole(models.ROLE_ADMIN, "System Administrator")
	tokenString, refreshToken, err := token.TokenGenerator(adminEmail, "Admin", "System", "", adminRole.Name)
	if err != nil {
		return fmt.Errorf("error generating tokens: %v", err)
	}
	// Tạo admin user
	firstName := "Admin"
	lastName := "System"
	password := string(hashedPassword)
	admin := models.User{
		First_Name:    &firstName,
		LastName:      &lastName,
		Email:         &adminEmail,
		Password:      &password,
		Role:          adminRole,
		Token:         &tokenString,
		Refresh_Token: &refreshToken,
		Create_At:     time.Now(),
		Update_At:     time.Now(),
	}

	// Insert admin user vào database
	_, err = UserCollection.InsertOne(ctx, admin)
	if err != nil {
		return fmt.Errorf("error creating admin user: %v", err)
	}
	log.Printf("Admin user created successfully with email: %s", adminEmail)
	return nil
}
