package transfer

import (
	"context"
	"reflect"
	"testing"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestRepository_List(t *testing.T) {
	type fields struct {
		storage map[string]vos.Transfer
	}

	type args struct {
		ctx       context.Context
		accountID string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []vos.Transfer
		wantErr bool
	}{
		{
			name: "should successfully return transfers list",
			fields: fields{
				storage: map[string]vos.Transfer{
					"123": {ID: "123", AccountOriginID: "1", AccountDestinationID: "2", Amount: 1000},
					"45":  {ID: "45", AccountOriginID: "1", AccountDestinationID: "3", Amount: 5000},
					"678": {ID: "678", AccountOriginID: "1", AccountDestinationID: "4", Amount: 10},
				},
			},
			args: args{
				ctx:       context.Background(),
				accountID: "1",
			},

			want: []vos.Transfer{
				{ID: "123", AccountOriginID: "1", AccountDestinationID: "2", Amount: 1000},
				{ID: "45", AccountOriginID: "1", AccountDestinationID: "3", Amount: 5000},
				{ID: "678", AccountOriginID: "1", AccountDestinationID: "4", Amount: 10},
			},
			wantErr: false,
		},
		{
			name: "should return successfully and empty list",
			fields: fields{
				storage: map[string]vos.Transfer{},
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				storage: tt.fields.storage,
			}
			got, err := r.List(tt.args.ctx, tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
