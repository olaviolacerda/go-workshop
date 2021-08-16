package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/olaviolacerda/account/internal/application"
	"github.com/olaviolacerda/account/internal/controller"
	"github.com/olaviolacerda/account/internal/infra"
	"time"

	chimid "github.com/go-chi/chi/v5/middleware"
)

func configureRoutes() *chi.Mux {
	router := chi.NewRouter()
	setMiddlewares(router)

	exchangeRepo := infra.NewExchangeRepository(infra.GetConfig("exchange_url"))
	accountRepo := infra.NewAccountRepository(infra.Connection())
	accountService := application.NewAccountService(exchangeRepo, accountRepo)
	accountController := controller.NewAccountController(accountService)

	router.Post("/v1/accounts", accountController.CreateAccount)
	router.Get("/v1/accounts/{account_id}/balance", accountController.GetBalance)
	return router
}

func setMiddlewares(mux *chi.Mux) {
	mux.Use(chimid.StripSlashes)
	mux.Use(chimid.Logger)
	mux.Use(chimid.Heartbeat("/liveness"))
	mux.Use(chimid.RequestID)
	mux.Use(chimid.Timeout(30 * time.Second))
}
