package domain

import (
	"testing"

	"github.com/matryer/is"
)

func TestCartItem_IncreaseFreeQuantity(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		cartItem := CartItem{
			SKU:          "1",
			FreeQuantity: 1,
		}

		res := cartItem.IncreaseFreeQuantity(1)
		is.Equal(2, res)
	})
}

func TestCartItem_IncreaseQuantity(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		cartItem := CartItem{
			SKU:      "1",
			Quantity: 1,
		}

		res := cartItem.IncreaseQuantity(1)
		is.Equal(2, res)
	})
}

func TestCartItem_TotalPrice(t *testing.T) {
	is := is.New(t)

	t.Run("success", func(t *testing.T) {
		cartItem := CartItem{
			SKU:             "1",
			Quantity:        3,
			Price:           30,
			DiscountedPrice: 20,
		}

		res := cartItem.TotalPrice()
		is.Equal(float64(60), res)
	})
}
