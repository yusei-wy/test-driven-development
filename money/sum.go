package money

var (
	_ Expression = (*Sum)(nil)
)

type Sum struct {
	augend Expression
	addend Expression
}

// Augend は Sum の加算元の Money を返します。
// 主にテストで使用されます。
func (s *Sum) Augend() *Money {
	return (s.augend).(*Money)
}

// Addend は Sum の加算先の Money を返します。
// 主にテストで使用されます。
func (s *Sum) Addend() *Money {
	return (s.addend).(*Money)
}

func NewSum(augend, addend Expression) *Sum {
	return &Sum{
		augend: augend,
		addend: addend,
	}
}

func (s *Sum) Plus(added Expression) Expression {
	return nil
}

// Reduce は augend と added を目的通貨に変換することで、通貨を一致させてから計算を行いその結果を返します。
func (s *Sum) Reduce(provider RateProvider, to Concurrency) *Money {
	amount := s.augend.Reduce(provider, to).Amount() +
		s.addend.Reduce(provider, to).Amount()
	return NewMoney(to, amount)
}
