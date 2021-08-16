package infra

import (
	"github.com/olaviolacerda/account/internal/domain"
	"gorm.io/gorm"
)

type accountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(conn *gorm.DB) domain.IAccountRepository {
	return accountRepository{conn}
}

func (ar accountRepository) Create(account domain.Account) error {
	if result := ar.DB.Create(account); result.Error != nil {
		return result.Error
	}
	return nil
}

func (ar accountRepository) FindById(id string) (domain.Account, error) {
	var account domain.Account

	result := ar.DB.Table("accounts").Find(&account, "id = ?", id)
	if result.Error != nil {
		return domain.Account{}, result.Error
	}

	return account, nil
}
