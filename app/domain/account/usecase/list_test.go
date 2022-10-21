package usecase

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/account"
	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestUsecase_List(t *testing.T) {
	type fields struct {
		repository account.Repository
	}

	tests := []struct {
		name    string
		fields  fields
		want    []vos.Account
		wantErr bool
	}{
		{
			name: "should successfully return account list",
			fields: fields{
				repository: account.RepositoryMock{
					AccountList: []vos.Account{
						{ID: "1", Name: "Joao", CPF: "22233344455", Secret: "Er5", CreatedAt: time.Date(2003, 10, 2, 0, 0, 0, 0, time.Local)},
						{ID: "2", Name: "Maria", CPF: "55533344455", Secret: "F35", CreatedAt: time.Date(1994, 16, 7, 0, 0, 0, 0, time.Local)},
						{ID: "3", Name: "Paula", CPF: "44455566647", Secret: "E43", CreatedAt: time.Date(1988, 1, 9, 0, 0, 0, 0, time.Local)},
					},
					GetAccountErr: nil,
					Account:       vos.Account{},
				},
			},
			want: []vos.Account{
				{ID: "1", Name: "Joao", CPF: "22233344455", Secret: "Er5", CreatedAt: time.Date(2003, 10, 2, 0, 0, 0, 0, time.Local)},
				{ID: "2", Name: "Maria", CPF: "55533344455", Secret: "F35", CreatedAt: time.Date(1994, 16, 7, 0, 0, 0, 0, time.Local)},
				{ID: "3", Name: "Paula", CPF: "44455566647", Secret: "E43", CreatedAt: time.Date(1988, 1, 9, 0, 0, 0, 0, time.Local)},
			},

			wantErr: false,
		},
		{
			name: "should successfully return an empty list",
			fields: fields{
				repository: account.RepositoryMock{
					AccountList: []vos.Account{},
					ListErr:     nil,
				},
			},
			want:    []vos.Account{},
			wantErr: false,
		},
		{
			name: "should return error when listing",
			fields: fields{
				repository: account.RepositoryMock{
					ListErr: errors.New("error when listing"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository: tt.fields.repository,
			}
			got, err := u.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
