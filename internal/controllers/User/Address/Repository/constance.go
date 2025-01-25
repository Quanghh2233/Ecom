package repository

import (
	"context"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddressRepository interface {
	AddAddress(ctx context.Context, userID primitive.ObjectID, address models.Address) error
	GetAddressCount(ctx context.Context, userID primitive.ObjectID) (int32, error)
	DeleteAddress(ctx context.Context, userID, addressID primitive.ObjectID) (int64, error)
	EditHomeAddress(ctx context.Context, userID primitive.ObjectID, address models.Address) (int64, error)
	EditWorkAddress(ctx context.Context, userID primitive.ObjectID, address models.Address) (int64, error)
}

type addressRepository struct {
	userCollection *mongo.Collection
}

func NewAddressRepository(userCollection *mongo.Collection) AddressRepository {
	return &addressRepository{userCollection: userCollection}
}
