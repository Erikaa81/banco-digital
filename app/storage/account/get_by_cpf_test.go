package account

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestRepository_GetByCPF(t *testing.T) {
	type fields struct {
		storage map[string]vos.Account
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
			name: "should successfully get account by CPF",
			fields: fields{

				storage: map[string]vos.Account{
					"22233344433": {ID: "1", Name: "Joao", CPF: "22233344433", CreatedAt: time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local)},
				},
			},
			args: args{
				ctx: context.Background(),
				cpf: "22233344433",
			},

			want:    vos.Account{ID: "1", Name: "Joao", CPF: "22233344433", CreatedAt: time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local)},
			wantErr: false,
		},
		{
			name: "should return error CPF not found",
			fields: fields{

				storage: map[string]vos.Account{},
			},
			args: args{
				ctx: context.Background(),
				cpf: "22233344433",
			},

			want:    vos.Account{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				storage: tt.fields.storage,
			}
			got, err := r.GetByCPF(tt.args.ctx, tt.args.cpf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetByCPF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetByCPF() = %v, want %v", got, tt.want)
			}
		})
	}
}
