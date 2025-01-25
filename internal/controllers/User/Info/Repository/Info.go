package repository

import (
	"context"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *userRepository) FindUserByID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	err := r.userCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindStoresByOwner(ctx context.Context, ownerID string) ([]models.Store, error) {
	cursor, err := r.storeCollection.Find(ctx, bson.M{"owner": ownerID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stores []models.Store
	if err = cursor.All(ctx, &stores); err != nil {
		return nil, err
	}
	return stores, nil
}

func (r *userRepository) CountStoresByOwner(ctx context.Context, ownerID string) (int64, error) {
	count, err := r.storeCollection.CountDocuments(ctx, bson.M{"owner": ownerID})
	if err != nil {
		return 0, err
	}
	return count, nil
}
