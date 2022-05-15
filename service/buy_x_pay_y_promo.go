package service

import (
	"github.com/milhamh95/checkr/domain"
)

type buyXPayYPromoCartStorage interface {
	Update(cartItem domain.CartItem) error
}

type buyXPayYPromoOffer struct {
	cartStorage buyXPayYPromoCartStorage
}

func NewBuyXPayYPromoOffer(cartStorage buyXPayYPromoCartStorage) *buyXPayYPromoOffer {
	return &buyXPayYPromoOffer{
		cartStorage: cartStorage,
	}
}

func (b *buyXPayYPromoOffer) Eligible(minQuantity, cartItemQuantity int) bool {
	return cartItemQuantity == minQuantity
}

func (b *buyXPayYPromoOffer) Apply(
	item domain.CartItem,
	product domain.Product,
	promo domain.Promo) error {
	if !b.Eligible(promo.BuyXPayY.MinQuantity, item.Quantity) {
		return nil
	}

	item.Quantity = 2
	item.FreeQuantity = 1
	err := b.cartStorage.Update(item)
	return err
}
