package money

type Expression interface {
	Times(multiplier int) Expression
	Plus(added Expression) Expression
	Reduce(provider RateProvider, to Concurrency) *Money
}

type RateProvider interface {
	Rate(from, to Concurrency) int
}
