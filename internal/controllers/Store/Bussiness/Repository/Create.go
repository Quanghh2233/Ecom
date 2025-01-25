package repository

import (
	"context"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *productRepository) CreateProduct(ctx context.Context, product *models.Product) (primitive.ObjectID, error) {
	product.Product_ID = primitive.NewObjectID()

	result, err := r.productCollection.InsertOne(ctx, product)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *productRepository) FindStoreByID(ctx context.Context, storeID primitive.ObjectID) (*models.Store, error) {
	var store models.Store
	err := r.storeCollection.FindOne(ctx, bson.M{"store_id": storeID}).Decode(&store)
	if err != nil {
		return nil, err
	}
	return &store, nil
}
