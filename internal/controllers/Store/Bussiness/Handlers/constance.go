package handlers

import service "github.com/Quanghh2233/Ecommerce/internal/controllers/Store/Bussiness/Service"

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}
