package repository

import (
	"context"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) (primitive.ObjectID, error)
	FindStoreByID(ctx context.Context, storeID primitive.ObjectID) (*models.Store, error)
	DeleteProduct(ctx context.Context, productID primitive.ObjectID) (int64, error)
	FindStoreByIDAndOwner(ctx context.Context, storeID, ownerID primitive.ObjectID) (*models.Store, error)
	ListProductsByStore(ctx context.Context, storeID primitive.ObjectID) ([]bson.M, error)
	UpdateProduct(ctx context.Context, productID primitive.ObjectID, update bson.M) (int64, int64, error)
	FindProductByID(ctx context.Context, productID primitive.ObjectID) (*models.Product, error)
}

type productRepository struct {
	productCollection *mongo.Collection
	storeCollection   *mongo.Collection
}

func NewProductRepository(productCollection, storeCollection *mongo.Collection) ProductRepository {
	return &productRepository{
		productCollection: productCollection,
		storeCollection:   storeCollection,
	}
}

func (r *productRepository) FindStoreByIDAndOwner(ctx context.Context, storeID, ownerID primitive.ObjectID) (*models.Store, error) {
	var store models.Store
	err := r.storeCollection.FindOne(ctx, bson.M{
		"store_id": storeID,
		"owner_id": ownerID,
	}).Decode(&store)
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *productRepository) FindProductByID(ctx context.Context, productID primitive.ObjectID) (*models.Product, error) {
	var product models.Product
	err := r.productCollection.FindOne(ctx, bson.M{"product_id": productID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
