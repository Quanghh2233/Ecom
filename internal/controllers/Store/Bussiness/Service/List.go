package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *productService) ListProductsByStore(ctx context.Context, storeID string) ([]bson.M, error) {
	objID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, err
	}

	return s.repo.ListProductsByStore(ctx, objID)
}
