package bank

import "tdd/money"

var _ money.RateProvider = (*Bank)(nil)

type Bank struct {
	// NOTE:
	// ポインタにしてしまうとアドレスが一致しないと値を取り出せないので、値型を指定する
	rates map[money.Pair]int
}

func NewBank() *Bank {
	return &Bank{
		rates: make(map[money.Pair]int),
	}
}

func (b *Bank) Reduce(source money.Expression, to money.Concurrency) *money.Money {
	return source.Reduce(b, to)
}

func (b *Bank) AddRate(from, to money.Concurrency, rate int) {
	b.rates[money.NewPair(from, to)] = rate
}

// Rate は指定された通貨への変換レートを返します。
func (b *Bank) Rate(from, to money.Concurrency) int {
	if rate, ok := b.rates[money.NewPair(from, to)]; ok {
		return rate
	}

	return 1
}
