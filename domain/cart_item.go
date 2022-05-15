package domain

type CartItem struct {
	SKU             string
	Name            string
	Price           float64
	DiscountedPrice float64
	Quantity        int
	FreeQuantity    int
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
