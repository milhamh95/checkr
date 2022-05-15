package domain

import (
	"github.com/matryer/is"
	"testing"
)

func TestProduct_CheckQuantity(t *testing.T) {
	is := is.New(t)
	t.Run("inventory product >= cart item quantity", func(t *testing.T) {
		product := Product{
			SKU:          "1",
			InventoryQty: 5,
		}

		res := product.CheckQuantity(2)
		is.Equal(true, res)
	})

	t.Run("inventory product < cart item quantity", func(t *testing.T) {
		product := Product{
			SKU:          "1",
			InventoryQty: 5,
		}

		res := product.CheckQuantity(10)
		is.Equal(false, res)
	})
}
