package account

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r Repository) GetBalance(ctx context.Context, id string) (int, error) {

	account, ok := r.storage[id]
	if !ok {
		return account.Balance, vos.ErrIDNotFound
	}
	return account.Balance, nil
}
