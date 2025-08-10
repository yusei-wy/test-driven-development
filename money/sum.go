package money

var _ Expression = (*Sum)(nil)

type Sum struct {
	augend *Money
	addend *Money
}

func (s *Sum) Augend() *Money {
	return s.augend
}

func (s *Sum) Addend() *Money {
	return s.addend
}

func NewSum(augend, addend *Money) *Sum {
	return &Sum{
		augend: augend,
		addend: addend,
	}
}

func (s *Sum) Reduce(to Concurrency) *Money {
	amount := s.augend.Amount() + s.addend.Amount()
	return NewMoney(to, amount)
}
