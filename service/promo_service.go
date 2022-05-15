package service

import (
	"errors"

	"github.com/milhamh95/checkr/domain"
)

type ApplyPromoOffer interface {
	Apply(cartItem domain.CartItem, product domain.Product, promo domain.Promo) error
}

type promoProductStorage interface {
	GetProduct(sku string) (domain.Product, error)
}

type promoStorageSource interface {
	GetPromo(code string) (domain.Promo, error)
}

type promoService struct {
	productStorage promoProductStorage
	promoStorage   promoStorageSource
	promoOffer     map[domain.PromoType]ApplyPromoOffer
}

func NewPromoService(
	productStorage promoProductStorage,
	promoStorage promoStorageSource,
	promoOffer map[domain.PromoType]ApplyPromoOffer,
) *promoService {
	return &promoService{
		productStorage: productStorage,
		promoStorage:   promoStorage,
		promoOffer:     promoOffer,
	}
}

func (p *promoService) Apply(cartItem domain.CartItem) error {
	product, err := p.productStorage.GetProduct(cartItem.SKU)
	if err != nil {
		return err
	}

	promo, err := p.promoStorage.GetPromo(product.Promo.Code)
	if err != nil {
		return err
	}

	promoOffer, ok := p.promoOffer[product.Promo.Type]
	if !ok {
		return errors.New("promo is not found")
	}

	err = promoOffer.Apply(cartItem, product, promo)
	return err

}
