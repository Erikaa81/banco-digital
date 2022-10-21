package account

import (
	"context"
	"testing"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestStore(t *testing.T) {
	ctx := context.Background()

	t.Run("should return successfully the stored account ", func(t *testing.T) {
		account := vos.Account{ID: "124", Name: "Jose", CPF: "77766655544"}
		r := Repository{
			storage: map[string]vos.Account{},
		}

		output, err := r.Store(ctx, account)
		if err != nil {
			t.Errorf("wanted error to be nil got: %v", err)
		}

		retrievedAccount, ok := r.storage[output.ID]
		if !ok {
			t.Errorf("I would like to find the account but can't find it")
		}

		if output != retrievedAccount {
			t.Errorf("accounts are the same and are not")
		}
	})
}
