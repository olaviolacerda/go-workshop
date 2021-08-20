package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/olaviolacerda/exchange/internal/api"
)

func TestExchangeController(t *testing.T) {
	app := api.New()

	t.Run("Get currencies", func(it *testing.T) {
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/exchange/currencies", nil)
		app.Server().Handler.ServeHTTP(res, req)

		if res.Result().StatusCode != 200 {
			it.Fatalf("Expected: status code 200 but get : %d", res.Result().StatusCode)
		}
	})

	t.Run("Get quotation", func(it *testing.T) {
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/exchange/quotation?to=brl&from=usd&val=10.0", nil)
		app.Server().Handler.ServeHTTP(res, req)

		if res.Result().StatusCode != 200 {
			it.Fatalf("Expected: status code 200 but get : %d", res.Result().StatusCode)
		}
	})

}
