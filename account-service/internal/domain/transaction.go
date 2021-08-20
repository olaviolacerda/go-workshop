package domain

import "time"

type Transaction struct {
	ID string
	AccountID string
	Currency string
	Operation int
	Value float64
	CreatedAt time.Time
}
