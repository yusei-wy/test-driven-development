package money_test

import (
	"tdd/money"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiplication(t *testing.T) {
	five := money.NewDollar(5)

	ten := five.Times(2)
	require.Equal(t, 10, ten.Amount())

	fifteen := five.Times(3)
	require.Equal(t, 15, fifteen.Amount())
}
