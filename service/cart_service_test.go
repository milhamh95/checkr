package service_test

import (
	"errors"
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

	t.Run("failed get product", func(t *testing.T) {
		fakeProductStorage := &counterfeiter.FakeCartProductStorage{}
		fakeProductStorage.GetProductReturns(domain.Product{}, errors.New("unknown err"))

		cartItem := domain.CartItem{
			SKU:             "1",
			Name:            "Laptop",
			Price:           30,
			DiscountedPrice: 30,
		}

		cartService := service.NewCartService(nil, fakeProductStorage, nil)

		cartItems := []domain.CartItem{
			cartItem,
		}
		err := cartService.AddItem(cartItems)
		is.Equal(errors.New("unknown err"), err)
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

		cartService := service.NewCartService(
			fakeCartStorage,
			fakeProductStorage,
			fakePromoservice,
		)

		cart, totalPrice, err := cartService.Calculate()
		is.NoErr(err)
		is.Equal(cartItems, cart)
		is.Equal(float64(200), totalPrice)
	})

	t.Run("empty cart", func(t *testing.T) {
		cartItems := []domain.CartItem{}
		fakeCartStorage := &counterfeiter.FakeCartStorageSource{}
		fakeCartStorage.FetchItemsReturns(cartItems)

		cartService := service.NewCartService(fakeCartStorage, nil, nil)

		_, _, err := cartService.Calculate()
		is.Equal(errors.New("cart is empty"), err)
	})

	t.Run("failed apply promo", func(t *testing.T) {
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
		fakePromoservice.ApplyReturns(errors.New("unknown error"))

		cartService := service.NewCartService(
			fakeCartStorage,
			nil,
			fakePromoservice,
		)

		_, _, err := cartService.Calculate()
		is.Equal(errors.New("unknown error"), err)
	})

	t.Run("failed get product", func(t *testing.T) {
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
		fakeProductStorage.GetProductReturns(domain.Product{}, errors.New("unknown err"))

		cartService := service.NewCartService(
			fakeCartStorage,
			fakeProductStorage,
			fakePromoservice,
		)

		_, _, err := cartService.Calculate()
		is.Equal(errors.New("unknown err"), err)
	})

	t.Run("failed to reduce inventory quantity", func(t *testing.T) {
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
		fakeProductStorage.ReduceInventoryQuantityReturns(errors.New("unknown err"))

		cartService := service.NewCartService(
			fakeCartStorage,
			fakeProductStorage,
			fakePromoservice,
		)

		_, _, err := cartService.Calculate()
		is.Equal(errors.New("unknown err"), err)
	})
}
