package mock

import (
	"errors"
	"time"

	"github.com/olaviolacerda/account/internal/domain"
)

type AccountRepositoryMock struct {
	mustFail bool
}

func NewAccountRepositoryMock(fail bool) AccountRepositoryMock {
	return AccountRepositoryMock{fail}
}

func (am AccountRepositoryMock) Create(data domain.Account) error {
	if am.mustFail {
		return errors.New("account creation failed")
	}
	return nil
}

func (am AccountRepositoryMock) FindById(id string) (domain.Account, error) {
	if am.mustFail {
		return domain.Account{}, errors.New("account not found")
	}
	return domain.Account{
		ID:        "73089a35-3b88-40df-ab27-64df5e58e343",
		Owner:     "c4c4c36a-f50b-4aca-9be0-f83c10e4d404",
		Balance:   1.0,
		Currency:  "BRL",
		CreatedAt: time.Now(),
	}, nil
}

func (am AccountRepositoryMock) UpdateBalance(id string, value float64) error {
	if am.mustFail {
		return errors.New("update account failed")
	}
	return nil
}

func (am AccountRepositoryMock) AddTransaction(data domain.Transaction) {

}
