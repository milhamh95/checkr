package storage

import (
	"fmt"

	"github.com/milhamh95/checkr/domain"
)

type promoStorage struct {
	items map[string]domain.Promo
}

func NewPromoStorage() *promoStorage {
	return &promoStorage{
		items: map[string]domain.Promo{
			"freeRaspberry": {
				Code: "freeRaspberry",
				Type: domain.BundlingPromo,
				Bundling: domain.Bundling{
					SKU: "234234",
				},
			},
			"3For2": {
				Code: "3For2",
				Type: domain.BuyXPayYPromo,
				BuyXPayY: domain.BuyXPayY{
					MinQuantity: 3,
					PayQuantity: 2,
				},
			},
			"10Discount": {
				Code: "10Discount",
				Type: domain.BulkDiscountPromo,
				BulkDiscount: domain.BulkDiscount{
					MinQuantity:        3,
					DiscountPercentage: 10,
				},
			},
		},
	}
}

func (p *promoStorage) GetPromo(code string) (domain.Promo, error) {
	promo, ok := p.items[code]
	if !ok {
		return domain.Promo{}, fmt.Errorf("promo with code %s is not found", code)
	}

	return promo, nil
}
