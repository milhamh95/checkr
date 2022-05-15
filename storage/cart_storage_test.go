package storage

import (
	"testing"

	"github.com/matryer/is"
	"github.com/milhamh95/checkr/domain"
)

func TestCartStorage_AddItem(t *testing.T) {
	item := domain.CartItem{
		SKU:      "1",
		Name:     "Laptop",
		Quantity: 1,
		Price:    30,
	}

	is := is.New(t)
	t.Run("add new item", func(t *testing.T) {
		cartStorage := NewCartStorage()
		cartStorage.AddItem(item)

		cartItem, err := cartStorage.GetItem("1")

		is.NoErr(err)
		is.Equal(item, cartItem)

	})

	t.Run("increase existing free quantity item", func(t *testing.T) {
		cartStorage := NewCartStorage()
		cartStorage.AddItem(item)
		cartStorage.AddItem(item)

		cartItem, err := cartStorage.GetItem("1")
		is.NoErr(err)

		is.Equal(2, cartItem.Quantity)

	})
}

func TestCartStorage_AddFreeItem(t *testing.T) {
	item := domain.CartItem{
		SKU:          "1",
		Name:         "Laptop",
		FreeQuantity: 1,
		Price:        30,
	}

	is := is.New(t)
	t.Run("add new free item", func(t *testing.T) {
		cartStorage := NewCartStorage()
		cartStorage.AddFreeItem(item)

		cartItem, err := cartStorage.GetItem("1")

		is.NoErr(err)
		is.Equal(item, cartItem)

	})

	t.Run("increase existing free quota item", func(t *testing.T) {
		cartStorage := NewCartStorage()
		cartStorage.AddFreeItem(item)
		cartStorage.AddFreeItem(item)

		cartItem, err := cartStorage.GetItem("1")
		is.NoErr(err)

		is.Equal(2, cartItem.FreeQuantity)

	})
}

func TestCartStorage_FetchItems(t *testing.T) {
	item := domain.CartItem{
		SKU:      "1",
		Name:     "Laptop",
		Quantity: 1,
		Price:    30,
	}

	is := is.New(t)
	t.Run("fetch items", func(t *testing.T) {
		cartStorage := NewCartStorage()
		cartStorage.AddItem(item)

		cartItems := cartStorage.FetchItems()
		is.Equal(1, len(cartItems))

	})
}

func TestCartStorage_GetItems(t *testing.T) {
	item := domain.CartItem{
		SKU:      "1",
		Name:     "Laptop",
		Quantity: 1,
		Price:    30,
	}

	is := is.New(t)
	t.Run("fetch items", func(t *testing.T) {
		cartStorage := NewCartStorage()
		cartStorage.items["1"] = item

		cartItem, err := cartStorage.GetItem("1")
		is.NoErr(err)
		is.Equal(item, cartItem)

	})
}

func TestCartStorage_UpdateItem(t *testing.T) {
	item := domain.CartItem{
		SKU:      "1",
		Name:     "Laptop",
		Quantity: 1,
		Price:    30,
	}

	is := is.New(t)
	t.Run("fetch items", func(t *testing.T) {
		cartStorage := NewCartStorage()
		cartStorage.items["1"] = item

		newItem := item
		newItem.Price = 50
		err := cartStorage.Update(newItem)
		is.NoErr(err)

		cartItem, err := cartStorage.GetItem("1")
		is.NoErr(err)
		is.Equal(newItem, cartItem)

	})
}

func TestCartStorage_RemoveCartItems(t *testing.T) {
	item := domain.CartItem{
		SKU:      "1",
		Name:     "Laptop",
		Quantity: 1,
		Price:    30,
	}

	is := is.New(t)
	t.Run("fetch items", func(t *testing.T) {
		cartStorage := NewCartStorage()
		cartStorage.items["1"] = item

		cartStorage.RemoveCartItems()

		res := cartStorage.FetchItems()
		is.Equal(0, len(res))
	})
}
