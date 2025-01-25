package handlers

import service "github.com/Quanghh2233/Ecommerce/internal/controllers/Store/Register/Service"

type StoreHandler struct {
	storeService service.StoreService
}

func NewStoreHandler(storeService service.StoreService) *StoreHandler {
	return &StoreHandler{storeService: storeService}
}
