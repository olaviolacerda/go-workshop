package domain

type IAccountRepository interface {
	Create(account Account) error
	FindById(id string) (Account, error)
}
