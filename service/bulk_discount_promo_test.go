package service_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/milhamh95/checkr/counterfeiter"
	"github.com/milhamh95/checkr/domain"
	"github.com/milhamh95/checkr/service"
)

func TestBulkDiscountPromoOffer_Apply(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		fakeCartStorage := &counterfeiter.FakeBulkDiscountCartStorage{}
		fakeCartStorage.UpdateReturns(nil)

		bulkDiscount := service.NewBulkDiscountPromoOffer(fakeCartStorage)

		cartItem := domain.CartItem{
			SKU:      "1",
			Quantity: 2,
		}

		product := domain.Product{
			SKU:   "1",
			Price: 30,
		}

		promo := domain.Promo{
			BulkDiscount: domain.BulkDiscount{
				MinQuantity:        2,
				DiscountPercentage: 10,
			},
		}

		err := bulkDiscount.Apply(cartItem, product, promo)
		is.NoErr(err)
	})
}

func TestBulkDiscountPromoOffer_Eligible(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		fakeCartStorage := &counterfeiter.FakeBulkDiscountCartStorage{}
		bulkDiscount := service.NewBulkDiscountPromoOffer(fakeCartStorage)

		cartItem := domain.CartItem{
			SKU:      "1",
			Quantity: 1,
		}

		product := domain.Product{
			SKU:   "1",
			Price: 30,
		}

		promo := domain.Promo{
			BulkDiscount: domain.BulkDiscount{
				MinQuantity:        2,
				DiscountPercentage: 10,
			},
		}

		err := bulkDiscount.Apply(cartItem, product, promo)
		is.NoErr(err)
	})
}
