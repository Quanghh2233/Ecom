package repository

import (
	"context"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRepository interface {
	AddProductToCart(ctx context.Context, productID primitive.ObjectID, userID string) error
	RemoveCartItem(ctx context.Context, productID primitive.ObjectID, userID string) error
	GetUserCart(ctx context.Context, userID primitive.ObjectID) (*models.User, error)
	AggregateCartTotal(ctx context.Context, userID primitive.ObjectID) (float64, error)
}

type cartRepository struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewCartRepository(prodCollection, userCollection *mongo.Collection) CartRepository {
	return &cartRepository{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}
