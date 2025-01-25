package service

import (
	"context"
)

func (s *userService) GetUserProfile(ctx context.Context, userID string) (map[string]interface{}, error) {
	user, err := s.repo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	stores, err := s.repo.FindStoresByOwner(ctx, userID)
	if err != nil {
		return nil, err
	}

	hasStores, err := s.repo.CountStoresByOwner(ctx, userID)
	if err != nil {
		return nil, err
	}

	profile := map[string]interface{}{
		"name":     *user.First_Name + " " + *user.LastName,
		"email":    user.Email,
		"phone":    user.Phone,
		"address":  user.Address_Details,
		"cart":     user.UserCart,
		"orders":   user.Order_Status,
		"stores":   stores,
		"hasStore": hasStores > 0,
	}

	return profile, nil
}
