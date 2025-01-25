package service

import (
	"context"
	"errors"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *productService) CreateProduct(ctx context.Context, userRole string, product *models.Product) (primitive.ObjectID, error) {
	if userRole != models.ROLE_ADMIN && userRole != models.ROLE_SELLER {
		return primitive.NilObjectID, errors.New("you do not have the required role to create a product")
	}

	if err := validateProduct(product); err != nil {
		return primitive.NilObjectID, err
	}

	store, err := s.repo.FindStoreByID(ctx, product.Store_ID)
	if err != nil {
		return primitive.NilObjectID, errors.New("store not found")
	}

	product.Store_Name = store.Name

	return s.repo.CreateProduct(ctx, product)
}

func validateProduct(product *models.Product) error {
	if product.Product_Name == "" {
		return errors.New("product name is required")
	}
	if product.Price == nil || *product.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if product.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if product.Store_ID.IsZero() {
		return errors.New("store_id is required")
	}
	return nil
}
