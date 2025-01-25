package service

import (
	"context"

	repository "github.com/Quanghh2233/Ecommerce/internal/controllers/Store/Bussiness/Repository"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService interface {
	CreateProduct(ctx context.Context, userRole string, product *models.Product) (primitive.ObjectID, error)
	DeleteProduct(ctx context.Context, userRole, userID, storeID, productID string) error
	ListProductsByStore(ctx context.Context, storeID string) ([]bson.M, error)
	UpdateProduct(ctx context.Context, userRole, userID, storeID, productID string, updateData *UpdateProductData) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}
