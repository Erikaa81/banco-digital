package account

import (
	"context"
	"fmt"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r *Repository) List(ctx context.Context) ([]vos.Account, error) {
	var accountsList []vos.Account
	var account vos.Account

	rows, err := r.Query(ctx, `SELECT id, name, cpf, secret, birthdate, created_at FROM accounts`)
	if err != nil {
		return []vos.Account{}, err
	}

	for rows.Next() {
		err := rows.Scan(
			&account.ID,
			&account.Name,
			&account.CPF,
			&account.Secret,
			&account.BirthDate,
			&account.CreatedAt,
		)
		if err != nil {
			return []vos.Account{}, err
		}
		accountsList = append(accountsList, account)
	}
	if len(accountsList) == 0 {
		return []vos.Account{}, fmt.Errorf("error while listing accounts: %w", err)
	}
	return accountsList, nil
}


// package account

// import (
// 	"context"
// 	"reflect"
// 	"testing"
// 	"time"
// 	d"github.com/erikaa81/banco-digital/app/domain/account/usecase"
// 	"github.com/erikaa81/banco-digital/app/domain/account"
// 	"github.com/erikaa81/banco-digital/app/domain/vos"
// )

// func TestRepository_Create(t *testing.T) {
// 	type fields struct {
// 		repository account.Repository
// 	}
// 	type args struct {
// 		ctx     context.Context
// 		account vos.CreateInput
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    vos.Account
// 		wantErr bool
// 	}{
// 		{
// 			name: "should return success when creating account",
// 			fields: fields{
// 				repository: account.RepositoryMock{
// 					CreateAccountOutput: vos.Account{
// 						ID:        "123",
// 						Name:      "Erika",
// 						CPF:       "22244455544",
// 						CreatedAt: time.Time{},
// 					},
// 					CreateAccountErr: nil,
// 				},
// 			},
// 			args: args{
// 				ctx: context.Background(),
// 				account: vos.CreateInput{
// 					Name:      "Erika",
// 					CPF:       "22244455544",
// 					Secret:    "iO9",
// 					BirthDate: time.Date(2000, 10, 2, 0, 0, 0, 0, time.Local),
// 				},
// 			},
// 			want: vos.Account{ID: "123",
// 				Name:      "Erika",
// 				CPF:       "22244455544",
// 				CreatedAt: time.Time{},
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := d.Usecase{
// 			}

// 			got, err := r.Create(tt.args.ctx, tt.args.account)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Repository.Create() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Repository.Create() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }