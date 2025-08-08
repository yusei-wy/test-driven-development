package money

import (
	"strconv"
	"tdd/expression"
)

type Concurrency string

const (
	Dollar Concurrency = "USD"
	Franc  Concurrency = "CHF"
)

var _ expression.Expression = (*Money)(nil)

type Money struct {
	currency Concurrency
	amount   int
}

func NewMoney(currency Concurrency, amount int) *Money {
	return &Money{
		currency: currency,
		amount:   amount,
	}
}

func NewDollar(amount int) *Money {
	return NewMoney(Dollar, amount)
}

func NewFranc(amount int) *Money {
	return NewMoney(Franc, amount)
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Currency() string {
	return string(m.currency)
}

func (m *Money) String() string {
	return strconv.Itoa(m.amount) + " " + m.Currency()
}

func (m *Money) Plus(added *Money) *Money {
	return NewMoney(m.currency, m.amount+added.amount)
}

func (m *Money) Times(multiplier int) *Money {
	return NewMoney(m.currency, m.amount*multiplier)
}

func (m *Money) Equals(other *Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}
