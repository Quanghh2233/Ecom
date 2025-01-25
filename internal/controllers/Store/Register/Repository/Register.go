package repository

import (
	"context"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *storeRepository) CheckStoreExists(ctx context.Context, email string) (bool, error) {
	count, err := r.storeCollection.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *storeRepository) RegisterStore(ctx context.Context, store *models.Store) (primitive.ObjectID, error) {
	store.Store_Id = primitive.NewObjectID()
	store.CreateAt = time.Now()
	store.Status = "active"

	result, err := r.storeCollection.InsertOne(ctx, store)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
