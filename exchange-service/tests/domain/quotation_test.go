package domain

import (
	"fmt"
	"testing"

	"github.com/olaviolacerda/exchange/internal/domain"
)

type ConversionScenario struct {
	Name   string
	From   string
	To     string
	Expect string
	Value  float64
}

func TestQuotationEntity(t *testing.T) {
	var from, to domain.Currency
	currency := domain.NewCurrency()

	// pattern: tests table
	suite := []ConversionScenario{
		{"Get quotation: EUR to BRL", "EUR", "BRL", "30.19", 5.0},
		{"Get quotation: EUR to USD", "EUR", "USD", "5.95", 5.0},
		{"Get quotation: USD to GBP", "USD", "GBP", "3.58", 5.0},
		{"Get quotation: GBP to EUR", "GBP", "EUR", "5.88", 5.0},
		{"Get quotation: GBP to BRL", "GBP", "BRL", "35.48", 5.0},
	}

	for _, sc := range suite {
		t.Run(sc.Name, func(it *testing.T) {
			currency.GetBySymbol(sc.From, &from)
			currency.GetBySymbol(sc.To, &to)

			quotation := domain.NewQuotation(to, from, sc.Value)
			quotation.Calculate()
			//fmt.Println(quotation)
			// converts to string due to precision truncate
			result := fmt.Sprintf("%.2f", quotation.Result)
			if result != sc.Expect {
				it.FailNow()
			}
		})
	}

}
