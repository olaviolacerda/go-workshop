package application

import (
	"github.com/olaviolacerda/exchange/internal/domain"
)

type ExchangeService struct{}

func NewExchangeService() ExchangeService {
	return ExchangeService{}
}

func (s ExchangeService) GetCurrencies(from, to string, value float64) (domain.Quotation, error) {
	var currency, currencyFrom, currencyTo domain.Currency

	currency = domain.NewCurrency()
	if err := currency.GetBySymbol(from, &currencyFrom); err != nil {
		return domain.Quotation{}, err
	}

	if err := currency.GetBySymbol(to, &currencyTo); err != nil {
		return domain.Quotation{}, err
	}

	quotation := domain.NewQuotation(currencyTo, currencyFrom, value)
	quotation.Calculate()

	return quotation, nil
}
