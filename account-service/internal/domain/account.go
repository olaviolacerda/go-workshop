package domain

import "time"

type Account struct {
	ID        string     `gorm:"column:id;size:36;primaryKey"`
	Owner     string     `gorm:"column:owner;size:36;not null"`
	Currency  string     `gorm:"column:currency;size:3;not null"`
	Balance   float64    `gorm:"column:balance"`
	CreatedAt time.Time  `gorm:"<-:create"`
	UpdatedAt *time.Time `gorm:"<-:update"`
}
