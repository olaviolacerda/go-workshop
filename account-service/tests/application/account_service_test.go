//+build integration

package application

import (
	"github.com/olaviolacerda/account/internal/application"
	"github.com/olaviolacerda/account/internal/infra"
	"os"
	"testing"
)

func TestSuiteAccountServiceInt(t *testing.T) {
	//change the current directory for test
	os.Chdir("../..")

	infra.InitDB()

	exchangeRepo := infra.NewExchangeRepository("http://localhost:7005")
	accountRepo := infra.NewAccountRepository(infra.Connection())
	service := application.NewAccountService(exchangeRepo, accountRepo)

	t.Run("Create account", func(it *testing.T) {
		fakeAccount := application.AccountRequest{
			Document: "31312321",
			Name:     "John Dow",
			Email:    "email@gmail.com",
			Currency: "USD",
			Age:      20,
		}

		if err := service.CreateAccount(fakeAccount); err == nil {
			it.Fatalf("fail: %v", err)
		}
	})
}
