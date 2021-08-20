package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/olaviolacerda/exchange/internal/application"
	"github.com/olaviolacerda/exchange/internal/domain"
)

func GetCurrencies(res http.ResponseWriter, req *http.Request) {
	currency := domain.NewCurrency()
	currencies := make([]CurrencyResponse, 0)

	for _, item := range currency.GetAll() {
		currencies = append(currencies, CurrencyResponse{item.Symbol, item.Name})
	}

	Json(res, currencies)
}

func GetQuotation(res http.ResponseWriter, req *http.Request) {
	from := strings.ToUpper(req.URL.Query().Get("from"))
	to := strings.ToUpper(req.URL.Query().Get("to"))
	value, err := strconv.ParseFloat(req.URL.Query().Get("val"), 32)
	if sendError(res, err, http.StatusBadRequest) {
		return
	}

	quotationSvc := application.NewExchangeService()
	quotation, err := quotationSvc.GetCurrencies(from, to, value)
	if sendError(res, err, http.StatusBadRequest) {
		return
	}

	log.Printf("[INFO] get quotation from %s to %s with %.2f", from, to, value)
	data := map[string]string{
		"from":  from,
		"to":    to,
		"value": fmt.Sprintf("%.2f", quotation.Result),
	}

	Json(res, data)
}

// Json formats the response to 'application/json' type
func Json(res http.ResponseWriter, data interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(data); sendError(res, err, http.StatusInternalServerError) {
		return
	}

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write(buf.Bytes())
}

func sendError(res http.ResponseWriter, err error, code int) bool {
	if err == nil {
		return false
	}

	http.Error(res, err.Error(), code)
	return true
}
