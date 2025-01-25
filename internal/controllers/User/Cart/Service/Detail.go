package service

import (
	"context"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *cartService) GetUserCart(ctx context.Context, userID string) (*models.User, float64, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, 0, err
	}

	user, err := s.repo.GetUserCart(ctx, userObjectID)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.AggregateCartTotal(ctx, userObjectID)
	if err != nil {
		return nil, 0, err
	}

	return user, total, nil
}
