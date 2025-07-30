package money_test

import (
	"tdd/money"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiplication(t *testing.T) {
	dollar := money.NewDollar(5)
	dollar.Times(2)

	require.Equal(t, 10, dollar.Amount())
}
