package service

import (
	"errors"
	"fmt"

	"github.com/milhamh95/checkr/domain"
)

type cartProductStorage interface {
	GetProduct(sku string) (domain.Product, error)
	ReduceInventoryQuantity(sku string, quantity int) error
}

type cartStorageSource interface {
	AddItem(cartItem domain.CartItem)
	FetchItems() []domain.CartItem
	RemoveCartItems()
	GetItem(sku string) (domain.CartItem, error)
}

type cartPromoService interface {
	Apply(cartItem domain.CartItem) error
}

type cartService struct {
	cartStorage    cartStorageSource
	productStorage cartProductStorage
	promoService   cartPromoService
}

func NewCartService(
	cartStorage cartStorageSource,
	productStorage cartProductStorage,
	promoService cartPromoService) *cartService {
	return &cartService{
		cartStorage:    cartStorage,
		productStorage: productStorage,
		promoService:   promoService,
	}
}

func (c *cartService) AddItem(cartItems []domain.CartItem) error {
	for _, cartItem := range cartItems {
		product, err := c.productStorage.GetProduct(cartItem.SKU)
		if err != nil {
			return err
		}

		totalCurrentCartItem := cartItem.Quantity
		currentCartItem, err := c.cartStorage.GetItem(cartItem.SKU)
		if err == nil {
			totalCurrentCartItem += currentCartItem.TotalQuantity()
		}

		if !product.CheckQuantity(totalCurrentCartItem) {
			err := fmt.Errorf("Item quantity for sku:%s is not enough.",
				cartItem.SKU,
			)
			return err
		}

		cartItem.Name = product.Name
		cartItem.Price = cartItem.RoundFloat(product.Price)
		cartItem.DiscountedPrice = cartItem.RoundFloat(product.Price)
		c.cartStorage.AddItem(cartItem)
	}

	return nil
}

func (c *cartService) Calculate() (cartItems []domain.CartItem, totalPrice float64, err error) {
	tmpCartItems := c.cartStorage.FetchItems()

	if len(tmpCartItems) == 0 {
		err = errors.New("cart is empty")
		return
	}

	for _, v := range tmpCartItems {
		err = c.promoService.Apply(v)
		if err != nil {
			return
		}

		err = c.reduceInventoryQuantity(v)
		if err != nil {
			return
		}

	}

	cartItems = c.cartStorage.FetchItems()
	for _, v := range cartItems {
		totalPrice += v.TotalPrice()
	}

	c.cartStorage.RemoveCartItems()

	return
}

func (c *cartService) reduceInventoryQuantity(
	cartItem domain.CartItem,
) error {
	product, err := c.productStorage.GetProduct(cartItem.SKU)
	if err != nil {
		return err
	}

	if !product.CheckQuantity(cartItem.TotalQuantity()) {
		return fmt.Errorf("item: %s is out of stock", product.Name)
	}

	err = c.productStorage.ReduceInventoryQuantity(cartItem.SKU, cartItem.TotalQuantity())
	return err
}
