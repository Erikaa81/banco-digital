package main

import (
	"fmt"
	"net/http"

	"github.com/erikaa81/banco-digital/app/domain/account/usecase"
	api "github.com/erikaa81/banco-digital/app/gateways/http"
	"github.com/erikaa81/banco-digital/app/gateways/http/account"
	storage "github.com/erikaa81/banco-digital/app/storage/account"
)

func main() {
	accountRepository := storage.NewRepository()
	accountUseCase := usecase.NewAccounteUseCase(accountRepository)
	accountHandler := account.NewHandler(accountUseCase)

	api := api.NewAPI(accountHandler)

	err := http.ListenAndServe(":8080", api.Router())
	if err != nil {
		fmt.Printf("Error initializing server: %s", err)
	}
}
