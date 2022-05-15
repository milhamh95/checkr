package roundfloat_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/milhamh95/checkr/pkg/roundfloat"
)

func Test_RoundFloat(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		is := is.New(t)

		res := roundfloat.RoundFloat(30.9890000)
		is.Equal(30.99, res)

	})
}
