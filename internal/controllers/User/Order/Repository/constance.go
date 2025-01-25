package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	InstantBuy(ctx context.Context, productID, userID primitive.ObjectID) error
	CancelOrder(ctx context.Context, userID, orderID primitive.ObjectID) (int64, error)
	CancelAllOrders(ctx context.Context, userID string) error
}

type orderRepository struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewOrderRepository(prodCollection, userCollection *mongo.Collection) OrderRepository {
	return &orderRepository{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}
