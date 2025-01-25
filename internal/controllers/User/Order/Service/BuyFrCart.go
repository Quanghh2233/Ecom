package service

import (
	"context"

	repository "github.com/Quanghh2233/Ecommerce/internal/controllers/User/Order/Repository"
)

type CartService interface {
	BuyItemsFromCart(ctx context.Context, userID string, selectedItems []string) ([]string, error)
}

type cartService struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) CartService {
	return &cartService{repo: repo}
}

func (s *cartService) BuyItemsFromCart(ctx context.Context, userID string, selectedItems []string) ([]string, error) {
	return s.repo.BuyItemsFromCart(ctx, userID, selectedItems)
}
