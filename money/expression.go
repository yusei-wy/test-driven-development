package money

type Expression interface {
	Plus(added Expression) Expression
	Reduce(provider RateProvider, to Concurrency) *Money
}

type RateProvider interface {
	Rate(from, to Concurrency) int
}
