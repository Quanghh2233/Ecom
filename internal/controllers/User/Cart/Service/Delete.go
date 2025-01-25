package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *cartService) RemoveCartItem(ctx context.Context, productID, userID string) error {
	productObjID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return err
	}

	return s.repo.RemoveCartItem(ctx, productObjID, userID)
}
