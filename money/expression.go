package money

type Expression interface {
	Reduce(to Concurrency) *Money
}
