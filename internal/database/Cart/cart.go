package cart

import (
	"errors"

	"github.com/Quanghh2233/Ecommerce/internal/models"
)

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("can't add this product to the cart")
	ErrCantRemoveItem     = errors.New("can't remove this item from the cart")
	ErrCantGetItem        = errors.New("was unable to get the item form the cart")
	ErrCantBuyCartItem    = errors.New("can't update the purchase")
	ErrCartEmpty          = errors.New("cart is empty")
	ErrCantCancelOrders   = errors.New("can't cancel orders")
	ErrCantBuyProduct     = errors.New("can't buy product")
	ErrorNoItemFound      = errors.New("no item found")
)

func calculateTotalPrice(cart []models.ProdutUser) int {
	var total int
	for _, item := range cart {
		total += item.Price
	}
	return total
}
