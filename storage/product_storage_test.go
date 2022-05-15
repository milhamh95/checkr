package storage

import (
	"github.com/matryer/is"
	"testing"
)

func TestProductStorage_Fetch(t *testing.T) {
	is := is.New(t)
	t.Run("success fetch items", func(t *testing.T) {
		productStorage := NewProductStorage()
		res := productStorage.Fetch()

		is.Equal(4, len(res))

	})
}

func TestProductStorage_GetProduct(t *testing.T) {
	is := is.New(t)

	t.Run("get product", func(t *testing.T) {
		productStorage := NewProductStorage()
		res, err := productStorage.GetProduct("234234")
		is.NoErr(err)
		is.Equal("234234", res.SKU)

	})
}

func TestProductStorage_ReduceInventoryQuantity(t *testing.T) {
	is := is.New(t)

	t.Run("get product", func(t *testing.T) {
		productStorage := NewProductStorage()
		err := productStorage.ReduceInventoryQuantity("234234", 1)
		is.NoErr(err)

		product, err := productStorage.GetProduct("234234")
		is.NoErr(err)
		is.Equal(1, product.InventoryQty)
	})
}
