package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/erikaa81/banco-digital/app/domain/account"
	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestUsecase_GetByID(t *testing.T) {
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
		want    vos.Account
		wantErr bool
	}{
		{
			name: "should successfully get account by ID",
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
				id:  "1233",
			},
			want: vos.Account{
				ID:   "1233",
				Name: "Joao",
				CPF:  "33344422233",
			},
			wantErr: false,
		},
		{
			name: "should return error when looking for ID at repository",
			fields: fields{
				repository: account.RepositoryMock{
					GetAccountErr: errors.New("error when looking for account"),
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "231",
			},
			want:    vos.Account{},
			wantErr: true,
		},
		{
			name: "should return error when empty ID is informed",
			args: args{
				ctx: context.Background(),
				id:  "",
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
