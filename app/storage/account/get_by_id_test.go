package account

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func TestRepository_GetByID(t *testing.T) {
	ctx := context.Background()

	t.Run("should successfully get account by ID", func(t *testing.T) {
		id := "1234"
		r := Repository{
			storage: map[string]vos.Account{
				id: {
					ID:        id,
					Name:      "Joao",
					CPF:       "55544433322",
					CreatedAt: time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local),
				},
			},
		}
		want := vos.Account{
			ID:        "1234",
			Name:      "Joao",
			CPF:       "55544433322",
			CreatedAt: time.Date(2022, 10, 16, 0, 0, 0, 0, time.Local),
		}
		got, err := r.GetByID(ctx, id)

		if want != got {
			t.Errorf("wanted error to be nil but got: %s", err)
		}

		if err != nil {
			t.Errorf("wanted error to be nil but got: %s", err)
		}
	})

	t.Run("should return error when account wasn't found with the given ID", func(t *testing.T) {
		r := Repository{
			storage: map[string]vos.Account{},
		}
		_, err := r.GetByID(ctx, "2341")
		if !errors.Is(err, vos.ErrIDNotFound) {
			t.Errorf("expected err %s, and received: %s", vos.ErrIDNotFound, err)
		}
	})
}
