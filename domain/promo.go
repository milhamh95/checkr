package domain

type Promo struct {
	Code         string
	Type         PromoType
	BulkDiscount BulkDiscount
	Bundling     Bundling
	BuyXPayY     BuyXPayY
}

type BulkDiscount struct {
	MinQuantity        int
	DiscountPercentage int
}

type Bundling struct {
	SKU string
}

type BuyXPayY struct {
	MinQuantity int
	PayQuantity int
}

type PromoType string

const (
	BundlingPromo     PromoType = "bundling"
	BulkDiscountPromo PromoType = "bulkDiscount"
	BuyXPayYPromo     PromoType = "buyXPayYPromo"
)
