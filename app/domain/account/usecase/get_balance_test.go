package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/erikaa81/banco-digital/app/domain/account"
	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestUsecase_GetBalance(t *testing.T) {
	type fields struct {
		repository account.Repository
	}

	type args struct {
		ctx context.Context
		id  string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "should successfully get balance by ID",
			fields: fields{
				repository: account.RepositoryMock{
					Account: vos.Account{
						ID: "1233",
						Balance: 100,
					},
					GetBalanceErr: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "1233",
			},
			want:100,
			wantErr: false,
		},
		{
			name: "should return error when looking for ID at repository",
			fields: fields{
				repository: account.RepositoryMock{
					Account: vos.Account{},
					GetBalanceErr: errors.New("error when looking for account"),
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "231",
			},
			want: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository: tt.fields.repository,
			}
			got, err := u.GetBalance(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
