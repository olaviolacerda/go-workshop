package domain

import (
	"fmt"
)

var (
	EUR = 1.00
	BRL = 6.038
	USD = 1.190
	GBP = 0.851
)

type Quotation struct {
	Source Currency
	Target Currency
	Value  float64
	Result float64
}

func (q Quotation) String() string {
	return fmt.Sprintf("[Source: %s, Target: %s, Value: %.2f, Result: %.2f]", q.Source.Symbol, q.Target.Symbol, q.Value, q.Result)
}

func NewQuotation(target, source Currency, val float64) Quotation {
	return Quotation{
		Source: source,
		Target: target,
		Value:  val,
		Result: 0.0}
}

func (q *Quotation) Calculate() {
	if q.Source.Symbol == "EUR" {
		q.Result = q.Value * currencyValue(q.Target.Symbol)
		return
	}

	eurValue := currencyValue("EUR") / currencyValue(q.Source.Symbol)
	q.Result = (eurValue * q.Value) * currencyValue(q.Target.Symbol)
}

func currencyValue(symbol string) float64 {
	switch symbol {
	case "BRL":
		return BRL
	case "EUR":
		return EUR
	case "GBP":
		return GBP
	case "USD":
		return USD
	}
	return 0.0
}
