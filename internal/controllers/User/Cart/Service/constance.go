package service

import (
	"context"

	repository "github.com/Quanghh2233/Ecommerce/internal/controllers/User/Cart/Repository"
	"github.com/Quanghh2233/Ecommerce/internal/models"
)

type CartService interface {
	AddProductToCart(ctx context.Context, productID, userID string) error
	RemoveCartItem(ctx context.Context, productID, userID string) error
	GetUserCart(ctx context.Context, userID string) (*models.User, float64, error)
}

type cartService struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) CartService {
	return &cartService{repo: repo}
}
