package repository

import (
	"context"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, userID string) (*models.User, error)
	FindStoresByOwner(ctx context.Context, ownerID string) ([]models.Store, error)
	CountStoresByOwner(ctx context.Context, ownerID string) (int64, error)
	UpdateUserInfo(ctx context.Context, userID string, update bson.M) (int64, error)
}

type userRepository struct {
	userCollection  *mongo.Collection
	storeCollection *mongo.Collection
}

func NewUserRepository(userCollection, storeCollection *mongo.Collection) UserRepository {
	return &userRepository{
		userCollection:  userCollection,
		storeCollection: storeCollection,
	}
}
