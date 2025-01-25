package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *productRepository) ListProductsByStore(ctx context.Context, storeID primitive.ObjectID) ([]bson.M, error) {
	cursor, err := r.productCollection.Find(ctx, bson.M{"store_id": storeID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []bson.M
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}
