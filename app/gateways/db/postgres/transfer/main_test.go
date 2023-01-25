package transfer

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/erikaa81/banco-digital/app/domain/vos"
	accountStorage "github.com/erikaa81/banco-digital/app/gateways/db/postgres/account"
)

func TestCreate(t *testing.T) {
	t.Run("should return success database integration", func(t *testing.T) {

		databaseUrl := "postgres://postgres:password@localhost:5777/clientes?sslmode=disable"

		dbConfig, err := pgxpool.ParseConfig(databaseUrl)
		if err != nil {
			log.Fatal(err)
		}
		conn, err := pgxpool.ConnectConfig(context.Background(), dbConfig)
		if err != nil {
			log.Fatal(err)
		}

		accountRepository := accountStorage.NewRepository(conn)
		ctx := context.Background()
		account1 := vos.Account{
			Name:      "Paulo",
			CPF:       "22244455544",
			BirthDate: time.Date(2000, 10, 2, 0, 0, 0, 0, time.UTC),
		}

		account2 := vos.Account{
			Name:      "Maria",
			CPF:       "44455566677",
			BirthDate: time.Date(2000, 10, 2, 0, 0, 0, 0, time.UTC),
		}

		_, err = conn.Exec(ctx, "DELETE FROM accounts")
		if err != nil {
			fmt.Println(err.Error())
		}

		accountOrigin, err := accountRepository.Create(ctx, account1)
		if err != nil {
			log.Fatal(err)
		}

		accountDestination, err := accountRepository.Create(ctx, account2)
		if err != nil {
			log.Fatal(err)
		}

		transferRepository := NewRepository(conn)
		input := vos.Transfer{
			AccountOriginID:      accountOrigin.ID,
			AccountDestinationID: accountDestination.ID,
			Amount:               100,
		}

		expectedtransfer := vos.Transfer{
			AccountOriginID:      accountOrigin.ID,
			AccountDestinationID: accountDestination.ID,
			Amount:               100,
		}

		_, err = conn.Exec(ctx, "DELETE FROM transfers")
		if err != nil {
			fmt.Println(err.Error())
		}

		transfer, err := transferRepository.Create(ctx, input)
		if err != nil {
			log.Fatal(err)
		}

		if transfer.ID == "" {
			t.Errorf("ID cannot be empty")
		}

		if !reflect.DeepEqual(transfer.AccountOriginID, expectedtransfer.AccountOriginID) {
			t.Errorf("Repository.Create() = %v, want %v", transfer.AccountOriginID, expectedtransfer.AccountOriginID)
		}
		if !reflect.DeepEqual(transfer.AccountDestinationID, expectedtransfer.AccountDestinationID) {
			t.Errorf("Repository.Create() = %v, want %v", transfer.AccountDestinationID, expectedtransfer.AccountDestinationID)
		}

		recoveredTransfer, err := transferRepository.GetByID(ctx, transfer.ID)
		if err != nil {
			t.Errorf("error = %v, wantErr ", err)
			return
		}

		expectedRecoveredTransfer := vos.Transfer{
			ID:                   transfer.ID,
			AccountOriginID:      transfer.AccountOriginID,
			AccountDestinationID: transfer.AccountDestinationID,
			Amount:               transfer.Amount,
			CreatedAt:            transfer.CreatedAt,
		}

		if recoveredTransfer != expectedRecoveredTransfer {
			t.Errorf("error")
		}

		list, err := transferRepository.List(ctx, accountDestination.ID)
		if err != nil {
			t.Errorf("Repository.Create() error = %v, wantErr ", err)
			return
		}
		expectedList := []vos.Transfer{
			{ID: transfer.ID, AccountOriginID: transfer.AccountOriginID, AccountDestinationID: transfer.AccountDestinationID, Amount: transfer.Amount, CreatedAt: transfer.CreatedAt},
		}
		if !reflect.DeepEqual(list, expectedList) {
			t.Errorf("Repository.Create() = %v, want %v", list, expectedList)
		}
	})
}
