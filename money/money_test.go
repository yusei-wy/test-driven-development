package money_test

import (
	"tdd/money"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiplication(t *testing.T) {
	five := money.NewDollar(5)
	require.True(t, money.NewDollar(10).Equals(five.Times(2)))
	require.True(t, money.NewDollar(15).Equals(five.Times(3)))
}

func TestEquality(t *testing.T) {
	require.True(t, money.NewDollar(5).Equals(money.NewDollar(5)))
	require.False(t, money.NewDollar(5).Equals(money.NewDollar(6)))
	require.False(t, money.NewFranc(5).Equals(money.NewDollar(5)))
}

func TestConcurrency(t *testing.T) {
	require.Equal(t, "USD", money.NewDollar(1).Currency())
	require.Equal(t, "CHF", money.NewFranc(1).Currency())
}
