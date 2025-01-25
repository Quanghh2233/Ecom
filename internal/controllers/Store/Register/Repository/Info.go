package repository

import (
	"context"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *storeRepository) GetStoreByID(ctx context.Context, storeID primitive.ObjectID) (*models.Store, error) {
	var store models.Store
	err := r.storeCollection.FindOne(ctx, bson.M{"store_id": storeID}).Decode(&store)
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *storeRepository) GetProductsByStoreID(ctx context.Context, storeID primitive.ObjectID, page, limit int) ([]bson.M, int64, error) {
	skip := (page - 1) * limit

	findOptions := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "create_at", Value: 1}})

	cursor, err := r.productCollection.Find(ctx, bson.M{"store_id": storeID}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var products []bson.M
	if err := cursor.All(ctx, &products); err != nil {
		return nil, 0, err
	}

	total, err := r.productCollection.CountDocuments(ctx, bson.M{"store_id": storeID})
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
