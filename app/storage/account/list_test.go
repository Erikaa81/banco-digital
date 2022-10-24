package account

import (
	"reflect"
	"testing"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestRepository_List(t *testing.T) {
	type fields struct {
		storage map[string]vos.Account
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
				storage: map[string]vos.Account{
					"123": {ID: "123", Name: "erika", CPF: "22233344455"},
					"45":  {ID: "45", Name: "Maria", CPF: "55533344455"},
					"678": {ID: "678", Name: "Paula", CPF: "44455566647"},
				},
			},
			want: []vos.Account{
				{ID: "123", Name: "erika", CPF: "22233344455"},
				{ID: "45", Name: "Maria", CPF: "55533344455"},
				{ID: "678", Name: "Paula", CPF: "44455566647"},
			},
			wantErr: false,
		},
		{
			name: "should return successfully an empty list",
			fields: fields{
				storage: map[string]vos.Account{},
			},
			want:  nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{
				storage: tt.fields.storage,
			}
			got, err := r.List()
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
