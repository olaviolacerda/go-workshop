package infra

import (
	"github.com/olaviolacerda/account/internal/common"
	"github.com/olaviolacerda/account/internal/domain"
	"log"
	"net/http"
	"time"
)

type exchangeRepository struct {
	exchangeUrl string
}

const timeout = 3 * time.Second

func NewExchangeRepository(url string) domain.IExchangeRepository {
	return exchangeRepository{url}
}

func (er exchangeRepository) GetCurrencies() []domain.Currency {
	// pointer to http.Client, reference to the struct
	httpClient := &http.Client{
		Timeout: timeout,
	}

	currencies := make([]domain.Currency, 0)
	response, err := httpClient.Get(er.exchangeUrl + "/v1/exchange/currencies")
	if err != nil {
		log.Printf("[ERROR] fail to get currencies. Reason: %s \n", err.Error())
		return currencies
	}

	if response.StatusCode != 200 {
		log.Printf("[ERROR] exchange service fail. Reason: %d \n", response.StatusCode)
		return currencies
	}

	if err := common.BindJSON(response.Body, &currencies); err != nil {
		log.Printf("[ERROR] fail to decode data from exchange. Reason: %v \n", err)
		return currencies
	}

	return currencies
}
