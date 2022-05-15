package storage

import (
	"fmt"

	"github.com/milhamh95/checkr/domain"
)

type productStorage struct {
	items map[string]domain.Product
}

func NewProductStorage() *productStorage {
	return &productStorage{
		items: map[string]domain.Product{
			"120P90": {
				SKU:          "120P90",
				Name:         "Google Home",
				Price:        49.99,
				InventoryQty: 10,
				Promo: domain.Promo{
					Code: "3For2",
					Type: domain.BuyXPayYPromo,
				},
			},
			"43N23P": {
				SKU:          "43N23P",
				Name:         "Macbook Pro",
				Price:        5399.99,
				InventoryQty: 5,
				Promo: domain.Promo{
					Code: "freeRaspberry",
					Type: domain.BundlingPromo,
				},
			},
			"A304SD": {
				SKU:          "A304SD",
				Name:         "Alexa Speaker",
				Price:        109.50,
				InventoryQty: 10,
				Promo: domain.Promo{
					Code: "10Discount",
					Type: domain.BulkDiscountPromo,
				},
			},
			"234234": {
				SKU:          "234234",
				Name:         "Raspberry Pi B",
				Price:        30,
				InventoryQty: 2,
			},
		},
	}
}

func (p *productStorage) GetProduct(sku string) (domain.Product, error) {
	product, ok := p.items[sku]
	if !ok {
		err := fmt.Errorf("product with sku: %s is not found", sku)
		return domain.Product{}, err
	}
	return product, nil
}

func (p *productStorage) Fetch() []domain.Product {
	products := []domain.Product{}
	for _, v := range p.items {
		products = append(products, v)
	}

	return products
}

func (p *productStorage) ReduceInventoryQuantity(sku string, quantity int) error {
	product, err := p.GetProduct(sku)
	if err != nil {
		return err
	}
	product.InventoryQty -= quantity
	p.items[sku] = product
	return nil
}
