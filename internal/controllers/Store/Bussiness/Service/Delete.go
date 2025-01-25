package service

import (
	"context"
	"errors"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *productService) DeleteProduct(ctx context.Context, userRole, userID, storeID, productID string) error {
	if userRole != models.ROLE_ADMIN && userRole != models.ROLE_SELLER {
		return errors.New("permission denied")
	}

	objStoreID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return errors.New("invalid store ID format")
	}

	if userRole == models.ROLE_SELLER {
		objUserID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			return errors.New("invalid user ID format")
		}

		_, err = s.repo.FindStoreByIDAndOwner(ctx, objStoreID, objUserID)
		if err != nil {
			return errors.New("you don't have permission to delete products from this store")
		}
	}

	objProductID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return errors.New("invalid product ID format")
	}

	deletedCount, err := s.repo.DeleteProduct(ctx, objProductID)
	if err != nil {
		return err
	}

	if deletedCount == 0 {
		return errors.New("product not found")
	}

	return nil
}
