package service

import (
	"context"

	repository "github.com/Quanghh2233/Ecommerce/internal/controllers/User/Cart/Repository"
	"go.mongodb.org/mongo-driver/bson"
)

type OrderService interface {
	GetUserOrders(ctx context.Context, userID string) ([]bson.M, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) GetUserOrders(ctx context.Context, userID string) ([]bson.M, error) {
	return s.repo.GetUserOrders(ctx, userID)
}
