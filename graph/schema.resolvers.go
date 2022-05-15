package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/milhamh95/checkr/domain"
	"github.com/milhamh95/checkr/graph/generated"
	"github.com/milhamh95/checkr/graph/model"
	"github.com/milhamh95/checkr/pkg/roundfloat"
)

const (
	MessageSuccess = "success"
)

func (r *mutationResolver) ScanItems(ctx context.Context, input model.ScanItems) (*model.ScanItemsResult, error) {
	if len(input.Items) == 0 {
		return nil, errors.New("cart item is empty")
	}

	cartItems := []domain.CartItem{}

	for _, v := range input.Items {
		cartItems = append(cartItems, domain.CartItem{
			SKU:      v.Sku,
			Quantity: v.Quantity,
		})
	}

	resp := &model.ScanItemsResult{}

	err := r.cartService.AddItem(cartItems)
	if err != nil {
		resp.Message = err.Error()
		return resp, nil
	}

	resp.Message = MessageSuccess
	return resp, nil
}

func (r *mutationResolver) Checkout(ctx context.Context) (*model.CheckoutResult, error) {
	resp := &model.CheckoutResult{}

	cartItems, totalPrice, err := r.cartService.Calculate()
	if err != nil {
		resp.Message = err.Error()
		return resp, nil
	}

	resp.TotalPrice = roundfloat.RoundFloat(totalPrice)
	resp.Message = MessageSuccess

	checkoutItems := []*model.CheckoutItem{}
	for _, v := range cartItems {
		checkoutItems = append(checkoutItems, &model.CheckoutItem{
			Sku:             v.SKU,
			Name:            v.Name,
			Price:           float64(v.Price),
			DiscountedPrice: float64(v.DiscountedPrice),
			Quantity:        v.Quantity,
			FreeQuantity:    v.FreeQuantity,
		})
	}

	resp.Items = checkoutItems
	return resp, nil
}

func (r *queryResolver) Echo(ctx context.Context) (string, error) {
	return "ok", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
