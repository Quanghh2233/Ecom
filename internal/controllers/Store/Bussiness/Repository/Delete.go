package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *productRepository) DeleteProduct(ctx context.Context, productID primitive.ObjectID) (int64, error) {
	result, err := r.productCollection.DeleteOne(ctx, bson.M{"product_id": productID})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
