package domain

import (
	"github.com/olaviolacerda/account/internal/application"
	"github.com/olaviolacerda/account/tests/mock"
	"testing"
)

func TestAccountServiceSuite(t *testing.T) {
	exchangeMock := mock.NewExchangeRepositoryMock()
	accountMock := mock.NewAccountRepositoryMock(false)

	accountService := application.NewAccountService(exchangeMock, accountMock)

	t.Run("Create new account", func(it *testing.T) {
		fakeAccount := application.AccountRequest{
			Document: "31312321",
			Name:     "John Dow",
			Email:    "email@gmail.com",
			Currency: "USD",
			Age:      20,
		}

		if err := accountService.CreateAccount(fakeAccount); err != nil {
			it.Fatalf("fail: %v", err)
		}
	})

	t.Run("Create new account with invalid currency", func(it *testing.T) {
		fakeAccount := application.AccountRequest{
			Document: "31312321",
			Name:     "John Dow",
			Email:    "email@gmail.com",
			Currency: "VND",
			Age:      20,
		}

		if err := accountService.CreateAccount(fakeAccount); err == nil {
			it.Fatalf("fail: %v", err)
		}
	})
}
