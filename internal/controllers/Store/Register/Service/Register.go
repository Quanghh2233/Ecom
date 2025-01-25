package service

import (
	"context"
	"errors"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *storeService) RegisterStore(ctx context.Context, userID string, store *models.Store) (primitive.ObjectID, error) {
	if err := validateStoreInput(store); err != nil {
		return primitive.NilObjectID, err
	}

	exists, err := s.repo.CheckStoreExists(ctx, store.Email)
	if err != nil {
		return primitive.NilObjectID, err
	}
	if exists {
		return primitive.NilObjectID, errors.New("store with this email already exists")
	}

	store.Owner = userID
	return s.repo.RegisterStore(ctx, store)
}

func validateStoreInput(store *models.Store) error {
	if store.Name == "" {
		return errors.New("store name is required")
	}
	if store.Email == "" {
		return errors.New("email is required")
	}
	if store.Phone == "" {
		return errors.New("phone number is required")
	}
	return nil
}
