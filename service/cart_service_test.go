package service_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/milhamh95/checkr/counterfeiter"
	"github.com/milhamh95/checkr/domain"
	"github.com/milhamh95/checkr/service"
)

func TestCartService_AddItem(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		fakeProductStorage := &counterfeiter.FakeCartProductStorage{}

		product := domain.Product{
			SKU:          "1",
			Name:         "Laptop",
			InventoryQty: 5,
			Price:        30,
		}
		fakeProductStorage.GetProductReturns(product, nil)

		cartItem := domain.CartItem{
			SKU:             "1",
			Name:            "Laptop",
			Price:           30,
			DiscountedPrice: 30,
		}
		fakeCartStorage := &counterfeiter.FakeCartStorageSource{}
		fakeCartStorage.AddItem(cartItem)
		fakeCartStorage.GetItemReturns(cartItem, nil)

		cartService := service.NewCartService(fakeCartStorage, fakeProductStorage, nil)

		cartItems := []domain.CartItem{
			cartItem,
		}
		err := cartService.AddItem(cartItems)
		is.NoErr(err)
	})
}

func TestCartService_Calculate(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		cartItems := []domain.CartItem{
			{
				SKU:             "1",
				Name:            "Laptop",
				Price:           100,
				DiscountedPrice: 100,
				Quantity:        2,
				FreeQuantity:    0,
			},
		}
		fakeCartStorage := &counterfeiter.FakeCartStorageSource{}
		fakeCartStorage.FetchItemsReturns(cartItems)

		fakePromoservice := &counterfeiter.FakeCartPromoService{}
		fakePromoservice.ApplyReturns(nil)

		fakeProductStorage := &counterfeiter.FakeCartProductStorage{}

		product := domain.Product{
			SKU:          "1",
			Name:         "Laptop",
			InventoryQty: 5,
		}
		fakeProductStorage.GetProductReturns(product, nil)
		fakeProductStorage.ReduceInventoryQuantityReturns(nil)

		cartService := service.NewCartService(fakeCartStorage, fakeProductStorage, fakePromoservice)

		cart, totalPrice, err := cartService.Calculate()
		if err != nil {
			is.NoErr(err)
			return
		}

		is.Equal(cartItems, cart)
		is.Equal(float64(200), totalPrice)
	})
}
