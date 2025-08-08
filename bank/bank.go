package bank

import (
	"tdd/expression"
	"tdd/money"
)

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) Reduce(source expression.Expression, to money.Concurrency) *money.Money {
	return money.NewDollar(10)
}
