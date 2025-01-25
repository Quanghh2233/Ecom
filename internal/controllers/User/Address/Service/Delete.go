package service

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *addressService) DeleteAddress(ctx context.Context, userID, addressID string) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	addressObjID, err := primitive.ObjectIDFromHex(addressID)
	if err != nil {
		return err
	}

	modifiedCount, err := s.repo.DeleteAddress(ctx, userObjID, addressObjID)
	if err != nil {
		return err
	}

	if modifiedCount == 0 {
		return errors.New("address not found or already deleted")
	}

	return nil
}
