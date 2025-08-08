package money_test

import (
	"tdd/bank"
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

func TestSimpleAddition(t *testing.T) {
	b := bank.NewBank()
	five := money.NewDollar(5)
	sum := five.Plus(five)
	reduced := b.Reduce(sum, money.Dollar)
	require.Equal(t, money.NewDollar(10), reduced)
}
