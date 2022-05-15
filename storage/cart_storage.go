package storage

import (
	"fmt"

	"github.com/milhamh95/checkr/domain"
)

type cartStorage struct {
	items map[string]domain.CartItem
}

func NewCartStorage() *cartStorage {
	return &cartStorage{
		items: map[string]domain.CartItem{},
	}
}

func (c *cartStorage) AddItem(cartItem domain.CartItem) {
	item, ok := c.items[cartItem.SKU]
	if ok {
		item.Quantity = item.IncreaseQuantity(cartItem.Quantity)
		c.items[cartItem.SKU] = item
		return
	}

	c.items[cartItem.SKU] = cartItem
}

func (c *cartStorage) AddFreeItem(cartItem domain.CartItem) {
	item, ok := c.items[cartItem.SKU]
	if ok {
		item.FreeQuantity = item.IncreaseFreeQuantity(cartItem.FreeQuantity)
		c.items[cartItem.SKU] = item
		return
	}

	c.items[cartItem.SKU] = cartItem
}

func (c *cartStorage) FetchItems() []domain.CartItem {
	cart := []domain.CartItem{}

	for _, v := range c.items {
		cart = append(cart, v)
	}

	return cart
}

func (c *cartStorage) GetItem(sku string) (domain.CartItem, error) {
	item, ok := c.items[sku]
	if !ok {
		return domain.CartItem{}, fmt.Errorf("cart item with sku:%s is not found", sku)
	}

	return item, nil
}

func (c *cartStorage) Update(cartItem domain.CartItem) error {
	_, err := c.GetItem(cartItem.SKU)
	if err != nil {
		return err
	}

	c.items[cartItem.SKU] = cartItem
	return nil
}

func (c *cartStorage) RemoveCartItems() {
	c.items = map[string]domain.CartItem{}
}
