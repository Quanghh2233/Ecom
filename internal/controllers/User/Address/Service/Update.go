package service

import (
	"context"
	"errors"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *addressService) EditHomeAddress(ctx context.Context, userID string, address models.Address) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	modifiedCount, err := s.repo.EditHomeAddress(ctx, userObjID, address)
	if err != nil {
		return err
	}

	if modifiedCount == 0 {
		return errors.New("home address not found")
	}

	return nil
}

func (s *addressService) EditWorkAddress(ctx context.Context, userID string, address models.Address) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	modifiedCount, err := s.repo.EditWorkAddress(ctx, userObjID, address)
	if err != nil {
		return err
	}

	if modifiedCount == 0 {
		return errors.New("work address not found")
	}

	return nil
}
