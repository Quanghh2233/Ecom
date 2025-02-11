package helper

import (
	"errors"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Role struct để lưu trong database

// Tạo role mới
func NewRole(name string, description string) (*models.Role, error) {
	if name == "" {
		name = models.DEFAULT_ROLE
	}

	switch name {
	case models.ROLE_ADMIN:
		return &models.Role{
			Role_ID:     primitive.NewObjectID(),
			Name:        models.ROLE_ADMIN,
			Description: description,
			Permissions: models.AdminPermissions,
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
		}, nil
	case models.ROLE_SELLER:
		return &models.Role{
			Role_ID:     primitive.NewObjectID(),
			Name:        models.ROLE_SELLER,
			Description: description,
			Permissions: models.SellerPermissions,
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
		}, nil
	case models.ROLE_BUYER:
		return &models.Role{
			Role_ID:     primitive.NewObjectID(),
			Name:        models.ROLE_BUYER,
			Description: description,
			Permissions: models.BuyerPermissions,
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
		}, nil
	default:
		return nil, errors.New("invalid role name")
	}
}

// Kiểm tra permission
