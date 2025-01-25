package service

import (
	"context"
	"errors"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *addressService) AddAddress(ctx context.Context, userID string, address models.Address) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	address.Address_id = primitive.NewObjectID()

	addressCount, err := s.repo.GetAddressCount(ctx, userObjID)
	if err != nil {
		return err
	}

	if addressCount >= 2 {
		return errors.New("address limit reached")
	}

	return s.repo.AddAddress(ctx, userObjID, address)
}
