package domain

import (
	"fmt"
)

type Currency struct {
	ID     string
	Symbol string
	Name   string
}

type CurrencyBuilder struct {
	Currency
}

func NewCurrency() Currency {
	return Currency{}
}

func (c Currency) IsValid() bool {
	for _, cur := range c.GetAll() {
		if c.Symbol != cur.Symbol && c.ID != cur.ID {
			return false
		}
	}
	return true
}

func (c Currency) GetBySymbol(symbol string, currency *Currency) error {
	for _, cur := range c.GetAll() {
		if symbol == cur.Symbol {
			currency.ID = cur.ID
			currency.Symbol = cur.Symbol
			currency.Name = cur.Name
			return nil
		}
	}
	return fmt.Errorf("%s not found", symbol)
}

func (c *CurrencyBuilder) WithID(value string) *CurrencyBuilder {
	c.ID = value
	return c
}

func (c *CurrencyBuilder) WithSymbol(value string) *CurrencyBuilder {
	c.Symbol = value
	return c
}

func (c *CurrencyBuilder) WithName(value string) *CurrencyBuilder {
	c.Name = value
	return c
}

func (c CurrencyBuilder) Build() Currency {
	return Currency{c.ID, c.Symbol, c.Name}
}

func (c Currency) GetAll() []Currency {
	builder := CurrencyBuilder{}
	return []Currency{
		builder.WithID("85214").WithSymbol("BRL").WithName("Real").Build(),
		builder.WithID("15439").WithSymbol("EUR").WithName("Euro").Build(),
		builder.WithID("65481").WithSymbol("GBP").WithName("Pounds").Build(),
		builder.WithID("99870").WithSymbol("USD").WithName("US Dolar").Build(),
	}
}
