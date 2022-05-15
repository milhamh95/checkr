package storage

import (
	"github.com/matryer/is"
	"testing"
)

func TestPromoStorage_GetPromo(t *testing.T) {
	is := is.New(t)
	t.Run("success get promo", func(t *testing.T) {
		promoStorage := NewPromoStorage()
		res, err := promoStorage.GetPromo("10Discount")
		is.NoErr(err)
		is.Equal("10Discount", res.Code)
	})
}
