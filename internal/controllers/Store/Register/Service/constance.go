package service

import (
	"context"

	repository "github.com/Quanghh2233/Ecommerce/internal/controllers/Store/Register/Repository"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoreService interface {
	GetStoreDetails(ctx context.Context, storeID string, page, limit int) (*StoreResponse, error)
	RegisterStore(ctx context.Context, userID string, store *models.Store) (primitive.ObjectID, error)
}

type storeService struct {
	repo repository.StoreRepository
}

func NewStoreService(repo repository.StoreRepository) StoreService {
	return &storeService{repo: repo}
}

type StoreResponse struct {
	Store      models.Store `json:"store"`
	Products   []bson.M     `json:"products"`
	Pagination struct {
		Total int64 `json:"total"`
		Page  int   `json:"page"`
		Limit int   `json:"limit"`
		Pages int64 `json:"pages"`
	} `json:"pagination"`
	Message string `json:"message"`
}
