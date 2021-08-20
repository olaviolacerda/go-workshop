package domain

type IExchangeRepository interface {
	GetCurrencies() []Currency
}
