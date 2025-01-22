package token

import (
	"fmt"
	"os"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/database"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Uid        string `json:"uid"`
	Role       string `json:"role"`
	TokenType  string `json:"token_type"`
	jwt.StandardClaims
}

var (
	UserData   *mongo.Collection = database.UserData(database.Client, "Users")
	SECRET_KEY                   = os.Getenv("SECRET_KEY")
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	SECRET_KEY = os.Getenv("SECRET_KEY")
}

func TokenGenerator(email, firstname, lastname, uid, role string) (string, string, error) {
	if SECRET_KEY == "" {
		return "", "", ErrSecretKeyMissing
	}

	claims := &SignedDetails{
		First_name: firstname,
		Last_name:  lastname,
		Email:      email,
		Uid:        uid,
		Role:       role,
		TokenType:  "access",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		// Email:     email,
		Uid:       uid,
		Role:      role,
		TokenType: "refresh",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(720 * time.Hour).Unix(),
		},
	}

	if role == models.ROLE_ADMIN {
		claims.StandardClaims.ExpiresAt = 0
		refreshClaims.StandardClaims.ExpiresAt = 0
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

var (
	ErrSecretKeyMissing = fmt.Errorf("SECRET_KEY is missing")
	ErrInvalidUserID    = fmt.Errorf("invalid user ID")
)
