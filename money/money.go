package money

type Money[T any] interface {
	Amount() int
	Times(multiplier int) T
	Equals(other T) bool
}

var _ Money[*Dollar] = (*Dollar)(nil)
var _ Money[*Franc] = (*Franc)(nil)

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}

func (d *Dollar) Amount() int {
	return d.amount
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{amount: d.amount * multiplier}
}

func (d *Dollar) Equals(other *Dollar) bool {
	return d.amount == other.amount
}

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{amount: amount}
}

func (f *Franc) Amount() int {
	return f.amount
}

func (f *Franc) Times(multiplier int) *Franc {
	return &Franc{amount: f.amount * multiplier}
}

func (f *Franc) Equals(other *Franc) bool {
	return f.amount == other.amount
}
