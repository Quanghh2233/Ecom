package service

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *orderService) CancelOrder(ctx context.Context, userID, orderID string) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	orderObjID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return err
	}

	modifiedCount, err := s.cancel.CancelOrder(ctx, userObjID, orderObjID)
	if err != nil {
		return err
	}

	if modifiedCount == 0 {
		return errors.New("order not found or already cancelled")
	}

	return nil
}

func (s *orderService) CancelAllOrders(ctx context.Context, userID string) error {
	return s.cancel.CancelAllOrders(ctx, userID)
}
