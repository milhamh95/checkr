package service

import (
	"log"

	"github.com/milhamh95/checkr/domain"
)

type bundlingPromoProductStorage interface {
	GetProduct(sku string) (domain.Product, error)
	ReduceInventoryQuantity(sku string, quantity int) error
}

type bundlingPromoCartStorage interface {
	AddFreeItem(cartItem domain.CartItem)
}

type bundlingPromoOffer struct {
	productStorage bundlingPromoProductStorage
	cartStorage    bundlingPromoCartStorage
}

func NewBundlingPromoOffer(
	productStorage bundlingPromoProductStorage,
	cartStorage bundlingPromoCartStorage) *bundlingPromoOffer {
	return &bundlingPromoOffer{
		productStorage: productStorage,
		cartStorage:    cartStorage,
	}
}

func (b *bundlingPromoOffer) Apply(
	item domain.CartItem,
	product domain.Product,
	promo domain.Promo,
) error {
	bundlingProduct, err := b.productStorage.GetProduct(promo.Bundling.SKU)
	if err != nil {
		return err
	}

	if bundlingProduct.InventoryQty == 0 {
		log.Printf("free product: %s is out of stock", bundlingProduct.Name)
		return nil
	}

	freeQuantity := item.Quantity
	if freeQuantity > bundlingProduct.InventoryQty {
		freeQuantity = bundlingProduct.InventoryQty
	}

	b.cartStorage.AddFreeItem(domain.NewCartItem(
		bundlingProduct.SKU,
		bundlingProduct.Name,
		bundlingProduct.Price,
		bundlingProduct.Price,
		0,
		freeQuantity,
	))

	err = b.productStorage.ReduceInventoryQuantity(item.SKU, freeQuantity)

	return err
}
