package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/olaviolacerda/account/internal/application"
	"github.com/olaviolacerda/account/internal/common"
	"net/http"
)

type accountController struct {
	accountService application.AccountService
}

func NewAccountController(accountService application.AccountService) accountController {
	return accountController{accountService}
}

func (a accountController) CreateAccount(res http.ResponseWriter, req *http.Request) {
	var data application.AccountRequest
	if err := common.BindJSON(req.Body, &data); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.accountService.CreateAccount(data); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	common.ToJSON(res, http.StatusCreated, map[string]string{
		"message": "Account has been created.",
	})
}

func (a accountController) GetBalance(res http.ResponseWriter, req *http.Request) {
	accountId := chi.URLParam(req, "account_id")
	balance, err := a.accountService.GetAccount(accountId)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	common.ToJSON(res, http.StatusOK, balance)
}
