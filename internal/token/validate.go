package token

import (
	"fmt"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	if SECRET_KEY == "" {
		msg = "SECRET_KEY is not set"
		return
	}

	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		if SECRET_KEY == "" {
			return nil, fmt.Errorf("SECRET_KEY is empty")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok || !token.Valid {
		msg = "Invalid token"
		return
	}

	isAdmin := claims.Role == models.ROLE_ADMIN

	if !isAdmin {
		now := time.Now().Unix()
		if claims.ExpiresAt < now {
			msg = "Token has expired"
			return nil, msg
		}
	}

	return claims, ""
}
