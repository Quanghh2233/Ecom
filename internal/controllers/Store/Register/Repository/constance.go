package repository

import (
	"context"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StoreRepository interface {
	GetStoreByID(ctx context.Context, storeID primitive.ObjectID) (*models.Store, error)
	GetProductsByStoreID(ctx context.Context, storeID primitive.ObjectID, page, limit int) ([]bson.M, int64, error)
	CheckStoreExists(ctx context.Context, email string) (bool, error)
	RegisterStore(ctx context.Context, store *models.Store) (primitive.ObjectID, error)
}

type storeRepository struct {
	storeCollection   *mongo.Collection
	productCollection *mongo.Collection
}

func NewStoreRepository(storeCollection, productCollection *mongo.Collection) StoreRepository {
	return &storeRepository{
		storeCollection:   storeCollection,
		productCollection: productCollection,
	}
}
