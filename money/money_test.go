package money_test

import (
	"tdd/money"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiplication(t *testing.T) {
	five := money.NewDollar(5)
	product := five.Times(2)

	require.Equal(t, 10, product.Amount())
}
