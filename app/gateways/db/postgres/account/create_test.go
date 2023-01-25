package account

// import (
// 	"context"
// 	"reflect"
// 	"testing"
// 	"time"

// 	"github.com/erikaa81/banco-digital/app/domain/vos"
// )

// func TestRepository_Create(t *testing.T) {
// 	type fields struct {
// 		repository FakeRepository
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
// 			args: args{
// 					ctx: context.Background(),
// 				account: vos.CreateInput{
// 					Name:      "Erika",
// 					CPF:       "22244455544",
// 					Secret:    "1q3",
// 					BirthDate: time.Date(2000, 10, 2, 0, 0, 0, 0, time.Local),
// 				},
// 			},

// 			fields: fields{repository: FakeRepository{QueryErr: nil}},
// 			want: vos.Account{
// 				ID:        "123",
// 				Name:      "Erika",
// 				CPF:       "22244455544",
// 				CreatedAt: time.Now(),
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := FakeRepository{
// 				QueryErr: nil,
// 			}

// 			got, err := r.QueryRow(tt.args.ctx, `name, cpf, secret, birthdate`, tt.args.account)
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Repository.Create() = %v, want %v", got, tt.want)
// 			}
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Repository.Create() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}

// 		})
// 	}
// }
