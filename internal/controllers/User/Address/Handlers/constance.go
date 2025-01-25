package handlers

import service "github.com/Quanghh2233/Ecommerce/internal/controllers/User/Address/Service"

type AddressHandler struct {
	addressService service.AddressService
}

func NewAddressHandler(addressService service.AddressService) *AddressHandler {
	return &AddressHandler{addressService: addressService}
}
