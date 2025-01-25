package service

import (
	"context"

	repository "github.com/Quanghh2233/Ecommerce/internal/controllers/User/Info/Repository"
)

type UserService interface {
	GetUserProfile(ctx context.Context, userID string) (map[string]interface{}, error)
	UpdateUserInfo(ctx context.Context, userID string, firstName, lastName, phone, address string) (int64, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewApplication(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
