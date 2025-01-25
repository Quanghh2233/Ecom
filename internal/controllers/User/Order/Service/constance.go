package service

import (
	"context"

	repository "github.com/Quanghh2233/Ecommerce/internal/controllers/User/Order/Repository"
)

type OrderService interface {
	CancelOrder(ctx context.Context, userID, orderID string) error
	CancelAllOrders(ctx context.Context, userID string) error
	InstantBuy(ctx context.Context, productID, userID string) error
}

type orderService struct {
	cancel  repository.OrderRepository
	instant repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		cancel:  repo,
		instant: repo,
	}
}
