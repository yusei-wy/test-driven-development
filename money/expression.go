package money

type Expression interface {
	Reduce(provider RateProvider, to Concurrency) *Money
}

type RateProvider interface {
	Rate(from, to Concurrency) int
}
