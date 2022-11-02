package transfer

import (
	"context"
	"reflect"
	"testing"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestRepository_Store(t *testing.T) {
	type fields struct {
		storage map[string]vos.Transfer
	}
	type args struct {
		ctx      context.Context
		transfer vos.Transfer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    vos.Transfer
		wantErr bool
	}{
		{
			name: "should return successfully the stored transfer",
			fields: fields{
				storage: map[string]vos.Transfer{},
			},
			args: args{
				ctx: nil,
				transfer: vos.Transfer{
					AccountOriginID:      "1",
					AccountDestinationID: "2",
					Amount:               10,
				},
			},
			want: vos.Transfer{
				AccountOriginID:      "1",
				AccountDestinationID: "2",
				Amount:               10,
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				storage: tt.fields.storage,
			}
			got, err := r.Store(tt.args.ctx, tt.args.transfer)
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

			tt.want.CreatedAt = got.CreatedAt

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.List() = %v, want %v", got, tt.want)
			}

			retrievedTransfer, ok := r.storage[got.ID]
			if !ok {
				t.Errorf("transfer not found")
			}

			if got != retrievedTransfer {
				t.Errorf("transfers are the same and are not")
			}
		})
	}
}
