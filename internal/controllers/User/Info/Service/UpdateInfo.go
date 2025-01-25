package service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *userService) UpdateUserInfo(ctx context.Context, userID string, firstName, lastName, phone, address string) (int64, error) {
	update := bson.M{
		"$set": bson.M{
			"first_name":      firstName,
			"last_name":       lastName,
			"phone":           phone,
			"address_details": address,
			"updated_at":      time.Now(),
		},
	}
	return s.repo.UpdateUserInfo(ctx, userID, update)
}
