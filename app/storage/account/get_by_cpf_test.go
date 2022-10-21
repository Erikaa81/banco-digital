package account

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestCpfExists(t *testing.T) {
	ctx := context.Background()

	t.Run("should successfully get account by CPF", func(t *testing.T) {
		cpf := "22233344433"
		r := Repository{
			storage: map[string]vos.Account{
				cpf: {
					ID:        "1",
					Name:      "Joao",
					CPF:       "22233344433",
					CreatedAt: time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local)},
			},
		}

		want := vos.Account{
			ID:        "1",
			Name:      "Joao",
			CPF:       "22233344433",
			CreatedAt: time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local),
		}

		got, err := r.GetByCPF(ctx, cpf)

		if want != got {
			t.Errorf("wanted error to be nil but got: %s", err)
		}

		if err != nil {
			t.Errorf("wanted error to be nil but got: %s", err)
		}
	})

	t.Run("should return error cpf not found", func(t *testing.T) {
		cpf := "22233344455"
		r := Repository{
			storage: map[string]vos.Account{},
		}

		_, err := r.GetByCPF(ctx, cpf)
		if !errors.Is(err, vos.ErrCPFNotFound) {
			t.Errorf("expected err %s, and received: %s", vos.ErrCPFNotFound, err)
		}
	})
}
