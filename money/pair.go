package money

type Pair struct {
	from Concurrency
	to   Concurrency
}

// NewPairt は指定された通貨のペアを作成します。
// Pair は map のキーとして使用されるので値型である必要があります。
func NewPair(from, to Concurrency) Pair {
	return Pair{
		from: from,
		to:   to,
	}
}

func (p Pair) Equals(other Pair) bool {
	return p.from == other.from && p.to == other.to
}

func (p Pair) HasCoee() int {
	return 0
}
