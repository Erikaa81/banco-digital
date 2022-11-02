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

func TestUsecase_GetByID(t *testing.T) {
	type fields struct {
		repository        transfer.Repository
		accountRepository account.Repository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    vos.Transfer
		wantErr bool
	}{
		{
			name: "should successfully get transfer by ID",
			fields: fields{
				repository: transfer.RepositoryMock{
					Transfer: vos.Transfer{
						ID:        "123",
						AccountOriginID: "1",
						AccountDestinationID: "2",
						Amount:    10,
						CreatedAt: time.Time{},
					},
					CreateTransferErr: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "123",
			},
			want: vos.Transfer{
				ID:        "123",
				AccountOriginID: "1",
				AccountDestinationID: "2",
				Amount:    10,
				CreatedAt: time.Time{},
			},
			wantErr: false,
		},
		{
			name: "should return error when looking for ID at repository",
			fields: fields{
				repository: transfer.RepositoryMock{
					GetTransferErr: errors.New("error when looking for transfer"),
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "231",
			},
			want:    vos.Transfer{},
			wantErr: true,
		},
		{
			name: "should return error when empty ID is informed",
			args: args{
				ctx: context.Background(),
				id:  "",
			},
			want:    vos.Transfer{},
			wantErr: true,
		},
		{
			name: "should successfully get transfer by ID",
			fields: fields{
				repository: transfer.RepositoryMock{
					Transfer: vos.Transfer{
						ID:        "123",
						AccountOriginID: "1",
						AccountDestinationID: "1",
						Amount:    10,
						CreatedAt: time.Time{},
					},
					CreateTransferErr: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "123",
			},
			want: vos.Transfer{
				ID:        "123",
				AccountOriginID: "1",
				AccountDestinationID: "1",
				Amount:    10,
				CreatedAt: time.Time{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository:        tt.fields.repository,
				accountRepository: tt.fields.accountRepository,
			}
			got, err := u.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
