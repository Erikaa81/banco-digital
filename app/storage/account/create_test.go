package account

import (
	"context"
	"testing"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestRepository_Store(t *testing.T) {
	type fields struct {
		storage map[string]vos.Account
	}
	type args struct {
		ctx     context.Context
		account vos.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    vos.Account
		wantErr bool
	}{
		{
			name: "should return successfully the stored account",
			fields: fields{
				storage: map[string]vos.Account{},
			},
			args: args{
				ctx: nil,
				account: vos.Account{
					ID:        "123",
					Name:      "Maria",
					CPF:       "2221133322",
					CreatedAt: time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local)},
			},
			want: vos.Account{
				ID:   "123",
				Name: "Maria",
				CPF:  "2221133322",
			},

			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				storage: tt.fields.storage,
			}
			got, err := r.Store(tt.args.ctx, tt.args.account)
			if err != nil {
				t.Errorf("wanted error to be nil got: %v", err)
			}

			if got.ID == "" {
				t.Errorf("wanted different than empty but got: %v", err)
			}
			tt.want.ID = got.ID

			if got.CreatedAt.IsZero() {
				t.Errorf("wanted different than empty but got: %v", err)
			}

			retrievedAccount, ok := r.storage[got.ID]
			if !ok {
				t.Errorf("account not found")
			}

			if got != retrievedAccount {
				t.Errorf("accounts are the same and are not")
			}
		})
	}
}
