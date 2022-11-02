package main

import (
	"fmt"
	"net/http"

	"github.com/erikaa81/banco-digital/app/domain/account/usecase"
	api "github.com/erikaa81/banco-digital/app/gateways/http"
	"github.com/erikaa81/banco-digital/app/gateways/http/account"
	storage "github.com/erikaa81/banco-digital/app/storage/account"
	"github.com/erikaa81/banco-digital/app/storage/transfer"
	usecaseTransfer"github.com/erikaa81/banco-digital/app/domain/transfer/usecase"
	httpTransfer"github.com/erikaa81/banco-digital/app/gateways/http/transfer"
)

func main() {
	accountRepository := storage.NewRepository()
	accountUseCase := usecase.NewAccounteUseCase(accountRepository)
	accountHandler := account.NewHandler(accountUseCase)

	transferRepository := transfer.NewRepository()
	transferUseCase := usecaseTransfer.NewTransferUseCase(transferRepository, accountRepository)
	transferHandler := httpTransfer.NewHandler(transferUseCase)

	api := api.NewAPI(accountHandler, transferHandler)

	err := http.ListenAndServe(":8080", api.Router())
	if err != nil {
		fmt.Printf("Error initializing server: %s", err)
	}
}
