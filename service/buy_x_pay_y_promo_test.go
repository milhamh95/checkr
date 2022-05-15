package service_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/milhamh95/checkr/counterfeiter"
	"github.com/milhamh95/checkr/domain"
	"github.com/milhamh95/checkr/service"
)

func TestBuyXPayYPromoOffer_Apply(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		fakeCartStorage := &counterfeiter.FakeBuyXPayYPromoCartStorage{}
		fakeCartStorage.UpdateReturns(nil)

		buyXPayY := service.NewBuyXPayYPromoOffer(fakeCartStorage)

		cartItem := domain.CartItem{
			SKU:      "1",
			Quantity: 3,
		}

		product := domain.Product{
			SKU:   "1",
			Price: 30,
		}

		promo := domain.Promo{
			BuyXPayY: domain.BuyXPayY{
				MinQuantity: 3,
				PayQuantity: 2,
			},
		}

		err := buyXPayY.Apply(cartItem, product, promo)
		is.NoErr(err)
	})
}

func TestBuyXPayYPromoOffer_Eligible(t *testing.T) {
	is := is.New(t)

	t.Run("is not eligible", func(t *testing.T) {
		buyXPayY := service.NewBuyXPayYPromoOffer(nil)

		res := buyXPayY.Eligible(3, 2)
		is.Equal(false, res)
	})
}
