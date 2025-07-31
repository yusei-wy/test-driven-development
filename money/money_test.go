package money_test

import (
	"tdd/money"
	"testing"

	"github.com/stretchr/testify/assert"
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
	require.True(t, money.NewFranc(5).Equals(money.NewFranc(5)))
	require.False(t, money.NewFranc(5).Equals(money.NewFranc(6)))
}

func TestFrancMultiplication(t *testing.T) {
	five := money.NewFranc(5)
	assert.True(t, money.NewFranc(10).Equals(five.Times(2)))
	assert.True(t, money.NewFranc(15).Equals(five.Times(3)))
}
