package bank

import (
	"tdd/money"
)

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) Reduce(source money.Expression, to money.Concurrency) *money.Money {
	return source.Reduce(to)
}
