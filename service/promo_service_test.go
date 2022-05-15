package service_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/milhamh95/checkr/counterfeiter"
	"github.com/milhamh95/checkr/domain"
	"github.com/milhamh95/checkr/service"
)

func TestPromoService_Apply(t *testing.T) {
	is := is.New(t)
	t.Run("success", func(t *testing.T) {
		fakeCartStorage := &counterfeiter.FakeBulkDiscountCartStorage{}
		promoOffer := map[domain.PromoType]service.ApplyPromoOffer{
			domain.BulkDiscountPromo: service.NewBulkDiscountPromoOffer(
				fakeCartStorage,
			),
		}

		fakeProductStorage := &counterfeiter.FakePromoProductStorage{}
		product := domain.Product{
			SKU:          "1",
			Name:         "Laptop",
			InventoryQty: 5,
			Promo: domain.Promo{
				Code: "abc",
				Type: domain.BulkDiscountPromo,
			},
		}
		fakeProductStorage.GetProductReturns(product, nil)

		fakePromoStorage := &counterfeiter.FakePromoStorageSource{}
		promo := domain.Promo{
			Code: "abc",
			Type: domain.BulkDiscountPromo,
			BulkDiscount: domain.BulkDiscount{
				MinQuantity:        2,
				DiscountPercentage: 10,
			},
		}
		fakePromoStorage.GetPromoReturns(promo, nil)

		promoService := service.NewPromoService(
			fakeProductStorage,
			fakePromoStorage,
			promoOffer,
		)

		cartItem := domain.CartItem{
			SKU:      "1",
			Name:     "Laptop",
			Price:    30,
			Quantity: 2,
		}
		err := promoService.Apply(cartItem)
		is.NoErr(err)
	})
}
