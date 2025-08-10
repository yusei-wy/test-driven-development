package bank_test

import (
	"tdd/bank"
	"tdd/money"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRate(t *testing.T) {
	b := bank.NewBank()
	b.AddRate(money.Franc, money.Dollar, 2)

	tests := []struct {
		name string
		from money.Concurrency
		to   money.Concurrency
		want int
	}{
		{
			name: "Dollar to Dollar (Default rate)",
			from: money.Dollar,
			to:   money.Franc,
			want: 1,
		},
		{
			name: "Franc to Franc (Default rate)",
			from: money.Franc,
			to:   money.Franc,
			want: 1,
		},
		{
			name: "Dollar to Franc (Default rate)",
			from: money.Dollar,
			to:   money.Franc,
			want: 1,
		},
		{
			name: "Franc to Dollar (Custom rate)",
			from: money.Franc,
			to:   money.Dollar,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := b.Rate(tt.from, tt.to)
			require.Equal(t, tt.want, got)
		})
	}
}
