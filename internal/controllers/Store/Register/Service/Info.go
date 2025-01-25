package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *storeService) GetStoreDetails(ctx context.Context, storeID string, page, limit int) (*StoreResponse, error) {
	objID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, err
	}

	store, err := s.repo.GetStoreByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	products, total, err := s.repo.GetProductsByStoreID(ctx, objID, page, limit)
	if err != nil {
		return nil, err
	}

	response := &StoreResponse{
		Store:    *store,
		Products: products,
		Message:  "Store retrieved successfully",
	}
	response.Pagination.Total = total
	response.Pagination.Page = page
	response.Pagination.Limit = limit
	response.Pagination.Pages = (total + int64(limit) - 1) / int64(limit)

	return response, nil
}
