package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/account"
	"github.com/erikaa81/banco-digital/app/domain/transfer"
	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestUsecase_Create(t *testing.T) {
	type fields struct {
		repository        transfer.Repository
		accountRepository account.Repository
	}

	type args struct {
		ctx   context.Context
		input vos.CreateTransferInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    vos.Transfer
		wantErr bool
	}{
		{
			name: "should return success when creating transfer",
			fields: fields{
				repository: transfer.RepositoryMock{
					TransferOutput:    vos.Transfer{ID: "123", AccountOriginID: "2", AccountDestinationID: "1", Amount: 10, CreatedAt: time.Time{}},
					CreateTransferErr: nil,
				},
				accountRepository: account.RepositoryMock{
					Account: vos.Account{
						ID: "1", Name: "Paula", CPF: "88877766655",
					},
				},
			},

			args: args{
				ctx: context.Background(),
				input: vos.CreateTransferInput{
					AccountOriginID:      "2",
					AccountDestinationID: "1",
					Amount:               10,
				},
			},
			want: vos.Transfer{
				ID:                   "123",
				AccountOriginID:      "2",
				AccountDestinationID: "1",
				Amount:               10,
				CreatedAt:            time.Time{},
			},
			wantErr: false,
		},

		{
			name: "should return error when creating transfer because account origin not found",
			fields: fields{
				accountRepository: account.RepositoryMock{
					GetAccountErr: vos.ErrAccountOriginNotFound,
				},
			},

			args: args{
				ctx: context.Background(),
				input: vos.CreateTransferInput{
					AccountOriginID: "1234",
					Amount:          100,
				},
			},
			want:    vos.Transfer{},
			wantErr: true,
		},
		{
			name: "should return error when creating transfer because account destination not found",
			fields: fields{
				accountRepository: account.RepositoryMock{
					GetAccountErr: vos.ErrAccountDestinationNotFound,
				},
			},

			args: args{
				ctx: context.Background(),
				input: vos.CreateTransferInput{
					AccountDestinationID: "1234",
					Amount:               10,
				},
			},
			want:    vos.Transfer{},
			wantErr: true,
		},
		{
			name: "should return error when persiting transfer fails",
			fields: fields{
				repository: transfer.RepositoryMock{
					TransferOutput:    vos.Transfer{},
					CreateTransferErr: errors.New("cannot persist transfer"),
				},
				accountRepository: account.RepositoryMock{
					Account: vos.Account{}},
			},

			args: args{
				ctx: context.Background(),
				input: vos.CreateTransferInput{
					AccountDestinationID: "765",
					Amount:               1000,
				},
			},
			want:    vos.Transfer{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository:        tt.fields.repository,
				accountRepository: tt.fields.accountRepository,
			}
			got, err := u.Create(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
