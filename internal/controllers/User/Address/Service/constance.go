package service

import (
	"context"

	repository "github.com/Quanghh2233/Ecommerce/internal/controllers/User/Address/Repository"
	"github.com/Quanghh2233/Ecommerce/internal/models"
)

type AddressService interface {
	AddAddress(ctx context.Context, userID string, address models.Address) error
	DeleteAddress(ctx context.Context, userID, addressID string) error
	EditHomeAddress(ctx context.Context, userID string, address models.Address) error
	EditWorkAddress(ctx context.Context, userID string, address models.Address) error
}

type addressService struct {
	repo repository.AddressRepository
}
