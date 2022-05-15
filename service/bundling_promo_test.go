package service_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/milhamh95/checkr/counterfeiter"
	"github.com/milhamh95/checkr/domain"
	"github.com/milhamh95/checkr/service"
)

func TestBundlingPromoOffer_Apply(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		fakeCartStorage := &counterfeiter.FakeBundlingPromoCartStorage{}

		fakeProductStorage := &counterfeiter.FakeBundlingPromoProductStorage{}

		product := domain.Product{
			SKU:   "2",
			Name:  "USB Drive",
			Price: 50,
		}
		fakeProductStorage.GetProductReturns(product, nil)

		bundling := service.NewBundlingPromoOffer(fakeProductStorage, fakeCartStorage)

		cartItem := domain.CartItem{
			SKU:      "1",
			Quantity: 2,
		}

		product1 := domain.Product{
			SKU:   "1",
			Price: 30,
		}

		promo := domain.Promo{
			Bundling: domain.Bundling{SKU: "2"},
		}

		err := bundling.Apply(cartItem, product1, promo)
		is.NoErr(err)

	})
}
