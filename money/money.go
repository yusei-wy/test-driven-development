package money

type moneyBase struct {
	amount int
}

func newMoneyBase(amount int) moneyBase {
	return moneyBase{amount: amount}
}

func (m *moneyBase) Amount() int {
	return m.amount
}

type Money[T any] interface {
	Times(multiplier int) T
	Equals(other T) bool
}

var _ Money[*Dollar] = (*Dollar)(nil)
var _ Money[*Franc] = (*Franc)(nil)

type Dollar struct {
	moneyBase
}

func NewDollar(amount int) *Dollar {
	return &Dollar{newMoneyBase(amount)}
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(d.amount * multiplier)
}

func (d *Dollar) Equals(other *Dollar) bool {
	return d.amount == other.amount
}

type Franc struct {
	moneyBase
}

func NewFranc(amount int) *Franc {
	return &Franc{newMoneyBase(amount)}
}

func (f *Franc) Times(multiplier int) *Franc {
	return NewFranc(f.amount * multiplier)
}

func (f *Franc) Equals(other *Franc) bool {
	return f.amount == other.amount
}
