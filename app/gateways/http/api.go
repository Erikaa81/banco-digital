package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/erikaa81/banco-digital/app/gateways/http/account"
	"github.com/erikaa81/banco-digital/app/gateways/http/healthcheck"
)

type API struct {
	HealthCheck *healthcheck.Handler
	Account     *account.Handler
}

func NewAPI(account *account.Handler) *API {
	return &API{
		HealthCheck: healthcheck.NewHandler(),
		Account:     account,
	}
}

func (a *API) Router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/health", a.HealthCheck.Get).Methods(http.MethodGet)
	router.HandleFunc("/accounts", a.Account.Create).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{id}", a.Account.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/accounts", a.Account.List).Methods(http.MethodGet)
	return router
}
