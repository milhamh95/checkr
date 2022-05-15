package domain

type Product struct {
	SKU          string
	Name         string
	Price        float64
	InventoryQty int
	Promo        Promo
}

func (p Product) CheckQuantity(quantity int) bool {
	return p.InventoryQty >= quantity
}
