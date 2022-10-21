package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/erikaa81/banco-digital/app/domain/account"
	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestUsecase_GetByCPF(t *testing.T) {
	type fields struct {
		repository account.Repository
	}

	type args struct {
		ctx context.Context
		cpf string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    vos.Account
		wantErr bool
	}{
		{
			name: "should successfully get account by CPF  ",
			fields: fields{
				repository: account.RepositoryMock{
					Account: vos.Account{
						ID:   "1233",
						Name: "Joao",
						CPF:  "33344422233",
					},
					CreateAccountErr: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				cpf: "33344422233",
			},
			want: vos.Account{
				ID:     "1233",
				Name:   "Joao",
				CPF:    "33344422233",
			},
			wantErr: false,
		},
		{
			name: "should return error when looking for CPF at repository",
			fields: fields{
				repository: account.RepositoryMock{
					GetAccountErr: errors.New("error when looking for account"),
				},
			},
			args: args{
				ctx: context.Background(),
				cpf: "23188877788",
			},
			want:    vos.Account{},
			wantErr: true,
		},
		{
			name: "should return error when empty CPF is informed",
			args: args{
				ctx: context.Background(),
				cpf: "",
			},
			want:    vos.Account{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository: tt.fields.repository,
			}
			got, err := u.GetByCPF(tt.args.ctx, tt.args.cpf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetByCPF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetByCPF() = %v, want %v", got, tt.want)
			}
		})
	}
}
