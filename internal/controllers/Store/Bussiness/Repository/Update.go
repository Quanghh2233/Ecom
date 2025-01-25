package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *productRepository) UpdateProduct(ctx context.Context, productID primitive.ObjectID, update bson.M) (int64, int64, error) {
	result, err := r.productCollection.UpdateOne(ctx, bson.M{"product_id": productID}, update)
	if err != nil {
		return 0, 0, err
	}
	return result.MatchedCount, result.ModifiedCount, nil
}
