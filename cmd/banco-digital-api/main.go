package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/erikaa81/banco-digital/app/domain/account/usecase"
	usecaseDeposit "github.com/erikaa81/banco-digital/app/domain/deposit/usecase"
	usecaseTransfer "github.com/erikaa81/banco-digital/app/domain/transfer/usecase"
	accountStorage "github.com/erikaa81/banco-digital/app/gateways/db/postgres/account"
	depositStorage "github.com/erikaa81/banco-digital/app/gateways/db/postgres/deposit"
	transferStorage "github.com/erikaa81/banco-digital/app/gateways/db/postgres/transfer"
	api "github.com/erikaa81/banco-digital/app/gateways/http"
	"github.com/erikaa81/banco-digital/app/gateways/http/account"
	httpDeposit "github.com/erikaa81/banco-digital/app/gateways/http/deposit"
	httpTransfer "github.com/erikaa81/banco-digital/app/gateways/http/transfer"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	ctx := context.Background()
	urlDB := "postgres://postgres:password@localhost:5778/digitalbanking"
	dbConfig, err := pgxpool.ParseConfig(urlDB)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgxpool.ConnectConfig(ctx, dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	accountRepository := accountStorage.NewRepository(conn)
	accountUseCase := usecase.NewAccounteUseCase(accountRepository)
	accountHandler := account.NewHandler(accountUseCase)

	transferRepository := transferStorage.NewRepository(conn)
	transferUseCase := usecaseTransfer.NewTransferUseCase(transferRepository, accountRepository)
	transferHandler := httpTransfer.NewHandler(transferUseCase)

	depositRepository := depositStorage.NewRepository(conn)
	depositUseCase := usecaseDeposit.NewDepositUseCase(depositRepository, accountRepository)
	depositHandler := httpDeposit.NewHandler(depositUseCase)

	api := api.NewAPI(accountHandler, transferHandler, depositHandler)

	err = http.ListenAndServe(":8080", api.Router())
	if err != nil {
		fmt.Printf("Error initializing server: %s", err)
	}
}
