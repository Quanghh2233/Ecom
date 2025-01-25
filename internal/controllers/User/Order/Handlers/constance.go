package handlers

import service "github.com/Quanghh2233/Ecommerce/internal/controllers/User/Order/Service"

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}
