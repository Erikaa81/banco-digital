package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/erikaa81/banco-digital/app/domain/account"
	"github.com/erikaa81/banco-digital/app/domain/transfer"
	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestUsecase_List(t *testing.T) {
	type fields struct {
		repository        transfer.Repository
		accountRepository account.Repository
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
			name: "should successfully return tranfers list",
			args: args{
				ctx:       context.Background(),
				accountID: "18a",
			},

			fields: fields{
				repository: transfer.RepositoryMock{
					TransferList: []vos.Transfer{
						{ID: "1", AccountOriginID: "18a", AccountDestinationID: "2", Amount: 200},
						{ID: "2", AccountOriginID: "18a", AccountDestinationID: "3", Amount: 150},
						{ID: "3", AccountOriginID: "18a", AccountDestinationID: "4", Amount: 300},
					},
					GetTransferErr: nil,
					Transfer:       vos.Transfer{},
				},
			},
			want: []vos.Transfer{
				{ID: "1", AccountOriginID: "18a", AccountDestinationID: "2", Amount: 200},
				{ID: "2", AccountOriginID: "18a", AccountDestinationID: "3", Amount: 150},
				{ID: "3", AccountOriginID: "18a", AccountDestinationID: "4", Amount: 300},
			},
			wantErr: false,
		},
		{
			name: "should successfully return an empty list",
			fields: fields{
				repository: transfer.RepositoryMock{
					TransferList: []vos.Transfer{},
					ListErr:      nil,
				},
			},
			want:    []vos.Transfer{},
			wantErr: false,
		},
		{
			name: "should return error when listing",
			fields: fields{
				repository: transfer.RepositoryMock{
					ListErr: errors.New("error when listing"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Usecase{
				repository:        tt.fields.repository,
				accountRepository: tt.fields.accountRepository,
			}
			got, err := u.List(tt.args.ctx, tt.args.accountID)
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
