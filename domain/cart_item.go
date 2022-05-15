package domain

import "math"

type CartItem struct {
	SKU             string
	Name            string
	Price           float64
	DiscountedPrice float64
	Quantity        int
	FreeQuantity    int
}

func (ci CartItem) RoundFloat(price float64) float64 {
	return math.Round(float64(price*100)) / 100
}

func NewCartItem(
	sku, name string,
	price, discountedPrice float64,
	quantity, freeQuantity int,
) CartItem {
	cartItem := CartItem{
		SKU:          sku,
		Name:         name,
		Quantity:     quantity,
		FreeQuantity: freeQuantity,
	}

	cartItem.Price = cartItem.RoundFloat(price)
	cartItem.DiscountedPrice = cartItem.RoundFloat(discountedPrice)
	return cartItem
}

func (ci CartItem) IncreaseQuantity(quantity int) int {
	return ci.Quantity + quantity
}

func (ci CartItem) IncreaseFreeQuantity(quantity int) int {
	return ci.FreeQuantity + quantity
}

func (ci CartItem) TotalPrice() float64 {
	return ci.DiscountedPrice * float64(ci.Quantity)
}

func (ci CartItem) TotalQuantity() int {
	return ci.Quantity + ci.FreeQuantity
}
