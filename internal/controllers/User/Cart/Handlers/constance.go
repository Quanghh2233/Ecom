package handlers

import service "github.com/Quanghh2233/Ecommerce/internal/controllers/User/Cart/Service"

type CartHandler struct {
	cartService service.CartService
}

func NewCartHandler(cartService service.CartService) *CartHandler {
	return &CartHandler{cartService: cartService}
}
