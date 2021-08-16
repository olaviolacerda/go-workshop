package api

import (
	"time"

	"github.com/go-chi/chi/v5"
	chimid "github.com/go-chi/chi/v5/middleware"

	"github.com/olaviolacerda/exchange/internal/controller"
)

func configureRoutes() *chi.Mux {
	router := chi.NewRouter()
	setMiddlewares(router)

	router.Get("/v1/exchange/currencies", controller.GetCurrencies)
	router.Get("/v1/exchange/quotation", controller.GetQuotation)

	return router
}

func setMiddlewares(mux *chi.Mux) {
	mux.Use(chimid.StripSlashes)
	mux.Use(chimid.Logger)
	mux.Use(chimid.Heartbeat("/liveness"))
	mux.Use(chimid.RequestID)
	mux.Use(chimid.Timeout(30 * time.Second))
}
