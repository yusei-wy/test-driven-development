package money

var (
	_ Expression = (*Sum)(nil)
)

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

// Reduce は augend と added を目的通貨に変換することで、通貨を一致させてから計算を行い、その結果を返します。
func (s *Sum) Reduce(provider RateProvider, to Concurrency) *Money {
	augend := s.augend.Reduce(provider, to)
	addend := s.addend.Reduce(provider, to)
	amount := augend.Amount() + addend.Amount()
	return NewMoney(to, amount)
}
