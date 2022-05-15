package service

import (
	"github.com/milhamh95/checkr/domain"
	"github.com/milhamh95/checkr/pkg/roundfloat"
)

type bulkDiscountCartStorage interface {
	Update(cartItem domain.CartItem) error
}

type bulkDiscountPromoOffer struct {
	cartStorage bulkDiscountCartStorage
}

func NewBulkDiscountPromoOffer(cartStorage bulkDiscountCartStorage) *bulkDiscountPromoOffer {
	return &bulkDiscountPromoOffer{
		cartStorage: cartStorage,
	}
}

func (b *bulkDiscountPromoOffer) Eligible(minQuantity, cartItemQuantity int) bool {
	return cartItemQuantity >= minQuantity
}

func (b *bulkDiscountPromoOffer) Apply(
	item domain.CartItem,
	product domain.Product,
	promo domain.Promo,
) error {
	if !b.Eligible(promo.BulkDiscount.MinQuantity, item.Quantity) {
		return nil
	}

	promoPriceReduction := product.Price *
		float64(promo.BulkDiscount.DiscountPercentage) / 100
	finalPrice := product.Price - promoPriceReduction

	item.DiscountedPrice = roundfloat.RoundFloat(finalPrice)
	err := b.cartStorage.Update(item)
	return err
}
