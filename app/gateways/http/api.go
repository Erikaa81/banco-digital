package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/erikaa81/banco-digital/app/gateways/http/account"
	"github.com/erikaa81/banco-digital/app/gateways/http/deposit"
	"github.com/erikaa81/banco-digital/app/gateways/http/healthcheck"
	"github.com/erikaa81/banco-digital/app/gateways/http/transfer"
)

type API struct {
	HealthCheck *healthcheck.Handler
	Account     *account.Handler
	Transfer    *transfer.Handler
	Deposit     *deposit.Handler
}

func NewAPI(account *account.Handler, transfer *transfer.Handler, deposit *deposit.Handler) *API {
	return &API{
		HealthCheck: healthcheck.NewHandler(),
		Account:     account,
		Transfer:    transfer,
		Deposit:     deposit,
	}
}

func (a *API) Router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/health", a.HealthCheck.Get).Methods(http.MethodGet)
	router.HandleFunc("/accounts", a.Account.Create).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{id}", a.Account.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{id}/balance", a.Account.GetBalance).Methods(http.MethodGet)
	router.HandleFunc("/accounts", a.Account.List).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{account-id}/transfers", a.Transfer.Create).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{account-id}/transfers", a.Transfer.List).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{account-id}/transfers/{transfer-id}", a.Transfer.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/deposits", a.Deposit.Create).Methods(http.MethodPost)

	return router
}
