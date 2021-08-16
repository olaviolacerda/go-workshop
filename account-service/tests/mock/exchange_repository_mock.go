package mock

import (
	"strconv"

	"github.com/olaviolacerda/account/internal/domain"
)

type ExchangeRepositoryMock struct {
}

func NewExchangeRepositoryMock() ExchangeRepositoryMock {
	return ExchangeRepositoryMock{}
}

func (er ExchangeRepositoryMock) GetCurrencies() []domain.Currency {
	return []domain.Currency{
		{"EUR", "Euro"},
		{"BRL", "Real"},
		{"USD", "US Dolar"},
	}
}

//func (er ExchangeRepositoryMock) GetQuotation(from, to string, value float64) domain.Quotation {
//	switch from {
//	case "eur":
//		{
//			if to == "brl" {
//				return domain.Quotation{From: from, To: to, Value: floatToStr(value * 6)}
//			}
//			if to == "usd" {
//				return domain.Quotation{From: from, To: to, Value: floatToStr(value * 0.90)}
//			}
//		}
//	case "brl":
//		{
//			if to == "eur" {
//				return domain.Quotation{From: from, To: to, Value: floatToStr(value * 0.16)}
//			}
//			if to == "usd" {
//				return domain.Quotation{From: from, To: to, Value: floatToStr(value * 0.6)}
//			}
//		}
//	case "usd":
//		{
//			if to == "eur" {
//				return domain.Quotation{From: from, To: to, Value: floatToStr(value * 1.20)}
//			}
//			if to == "brl" {
//				return domain.Quotation{From: from, To: to, Value: floatToStr(value * 0.20)}
//			}
//		}
//	}
//
//	return domain.Quotation{}
//}

func floatToStr(value float64) string {
	return strconv.FormatFloat(value, '.', 2, 64)
}
