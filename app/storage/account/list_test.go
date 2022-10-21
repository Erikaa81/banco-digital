package account

import (
	"testing"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestList(t *testing.T) {
	t.Run("should successfully return account list", func(t *testing.T) {
		account1 := vos.Account{ID: "123", Name: "erika", CPF: "22233344455"}
		account2 := vos.Account{ID: "45", Name: "Maria", CPF: "55533344455"}
		account3 := vos.Account{ID: "678", Name: "Paula", CPF: "44455566647"}

		r := Repository{
			storage: map[string]vos.Account{},
		}

		want := []vos.Account{account1, account2, account3}
		got, err := r.List()
		for i, v := range got {
			if want[i] != v {
				t.Errorf("want:%+v got%+v", want[i], got[i])
			}
		}

		if err != nil {
			t.Errorf("wanted error to be nil and got: %v", err)
		}
	})

	t.Run("should return error and an empty list when listing", func(t *testing.T) {
		r := Repository{
			storage: nil,
		}
		got, err := r.List()
		if len(got) != 0 {
			t.Errorf("wanted empty list but got: %+v", got)
		}

		if err != nil {
			t.Errorf("wanted error to be nil and got: %v", err)
		}
	})
}
