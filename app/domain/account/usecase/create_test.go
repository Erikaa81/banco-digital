package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/account"
	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestUsecase_Create(t *testing.T) {
	type fields struct {
		repository account.Repository
	}
	type args struct {
		ctx   context.Context
		input vos.CreateInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    vos.Account
		wantErr bool
	}{
		{
			name: "should return success when creating account",
			fields: fields{
				repository: account.RepositoryMock{
					CreateAccountOutput: vos.Account{
						ID:        "123",
						Name:      "Erika",
						CPF:       "22244455544",
						CreatedAt: time.Time{},
					},
					CreateAccountErr: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				input: vos.CreateInput{
					Name:      "Erika",
					CPF:       "22244455544",
					Secret:    "iO9",
					BirthDate: time.Date(2000, 10, 2, 0, 0, 0, 0, time.Local),
				},
			},
			want: vos.Account{ID: "123",
				Name:      "Erika",
				CPF:       "22244455544",
				CreatedAt: time.Time{},
			},
			wantErr: false,
		},
		{
			name: "should return error when creating account because client not of legal age",
			fields: fields{
				repository: account.RepositoryMock{
					CreateAccountOutput: vos.Account{},
					CreateAccountErr:    ErrInvalidAge,
				},
			},

			args: args{
				ctx: context.Background(),
				input: vos.CreateInput{
					Name:      "Joao",
					CPF:       "33355566677",
					Secret:    "y7d",
					BirthDate: time.Date(2007, 10, 2, 0, 0, 0, 0, time.Local),
				},
			},
			want:    vos.Account{},
			wantErr: true,
		},
		{
			name: "should return error when creating account because already an account for this CPF",
			fields: fields{
				repository: account.RepositoryMock{
					GetAccountErr: vos.ErrAccountNotFound,
				},
			},

			args: args{
				ctx: context.Background(),
				input: vos.CreateInput{
					Name:      "Joao",
					CPF:       "33355566677",
					Secret:    "y7d",
					BirthDate: time.Date(2007, 10, 2, 0, 0, 0, 0, time.Local),
				},
			},
			want:    vos.Account{},
			wantErr: true,
		},
		{
			name: "should return error when persiting account fails",
			fields: fields{
				repository: account.RepositoryMock{
					CreateAccountOutput: vos.Account{},
					CreateAccountErr:    errors.New("cannot persist account"),
				},
			},

			args: args{
				ctx: context.Background(),
				input: vos.CreateInput{
					Name:      "Erika",
					CPF:       "22244455544",
					Secret:    "iO9",
					BirthDate: time.Date(2003, 10, 2, 0, 0, 0, 0, time.Local),
				},
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
