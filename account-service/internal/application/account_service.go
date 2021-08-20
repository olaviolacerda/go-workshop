package application

import (
	"errors"
	"github.com/olaviolacerda/account/internal/common"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/olaviolacerda/account/internal/domain"
)

type AccountService struct {
	exchangeRepository domain.IExchangeRepository
	accountRepository  domain.IAccountRepository
}

var (
	validate = validator.New()
)

func NewAccountService(exchangeRepo domain.IExchangeRepository, accountRepo domain.IAccountRepository) AccountService {
	return AccountService{exchangeRepo, accountRepo}
}

func (as AccountService) CreateAccount(data AccountRequest) error {
	if err := validate.Struct(data); err != nil {
		return err
	}

	log.Printf("[DEBUG] info: %v", data)

	if !as.isValidCurrency(data.Currency) {
		return errors.New("invalid currency")
	}

	generatedId, _ := uuid.NewRandom()
	account := domain.Account{
		ID:        generatedId.String(),
		Owner:     data.Document,
		Currency:  data.Currency,
		Balance:   0.0,
		CreatedAt: time.Now(),
	}

	if err := as.accountRepository.Create(account); err != nil {
		return err
	}
	return nil
}

func retrieveLastUpdated(account domain.Account) string {
	var lastUpdated time.Time
	lastUpdated = *account.UpdatedAt

	if account.UpdatedAt == nil {
		lastUpdated = account.CreatedAt
	}

	return common.DateToStr(lastUpdated)
}

func (as AccountService) GetAccount(id string) (AccountBalanceResponse, error) {
	var account domain.Account
	var err error

	if account, err = as.accountRepository.FindById(id); err != nil {
		return AccountBalanceResponse{}, err
	}

	accountBalanceResponse := AccountBalanceResponse{
		Account:    account.ID,
		Balance:    account.Balance,
		LastUpdate: retrieveNonEmptyUpdatedAt(account),
	}

	return accountBalanceResponse, nil
}

func (as AccountService) isValidCurrency(value string) bool {
	data := as.exchangeRepository.GetCurrencies()
	for _, currency := range data {
		if currency.Symbol == value {
			return true
		}
	}
	return false
}
