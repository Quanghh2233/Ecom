package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *orderService) InstantBuy(ctx context.Context, productID, userID string) error {
	productObjID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return err
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	return s.instant.InstantBuy(ctx, productObjID, userObjID)
}
