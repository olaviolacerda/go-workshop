package domain

import (
	"testing"

	"github.com/olaviolacerda/exchange/internal/domain"
)

func TestCurrencyEntity(t *testing.T) {
	var curr domain.Currency

	t.Run("Get all currencies", func(it *testing.T) {
		currency := domain.NewCurrency()
		list := currency.GetAll()
		if len(list) != 4 {
			it.FailNow()
		}

	})

	t.Run("Find currency by symbol: BRL", func(it *testing.T) {
		currency := domain.NewCurrency()
		if err := currency.GetBySymbol("BRL", &curr); err != nil {
			it.FailNow()
		}
	})

	t.Run("Check invalid currency", func(it *testing.T) {
		builder := domain.CurrencyBuilder{}
		currency := builder.WithID("85690").WithSymbol("CAD").WithName("Canadian Dolar").Build()
		if currency.IsValid() {
			it.FailNow()
		}
	})
}
