package graph

import "github.com/milhamh95/checkr/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type checkoutCartService interface {
	AddItem(cartItems []domain.CartItem) error
	Calculate() (cartItems []domain.CartItem, totalPrice float64, err error)
}

func NewResolver(cartService checkoutCartService) *Resolver {
	return &Resolver{
		cartService: cartService,
	}
}

type Resolver struct {
	cartService checkoutCartService
}
