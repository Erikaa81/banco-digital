package account

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestRepository_GetByID(t *testing.T) {
	type fields struct {
		storage map[string]vos.Account
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

				storage: map[string]vos.Account{
					"22233344433": {ID: "1", Name: "Joao", CPF: "22233344433", CreatedAt: time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local)},
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "1",
			},

			want:    vos.Account{ID: "1", Name: "Joao", CPF: "22233344433", CreatedAt: time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local)},
			wantErr: false,
		},
		{
			name: "should return error ID not found",
			fields: fields{

				storage: map[string]vos.Account{},
			},
			args: args{
				ctx: context.Background(),
				id:  "1",
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
			got, err := r.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
