package account

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r Repository) GetByID(ctx context.Context, id string) (vos.Account, error) {
	account, ok := r.storage[id]
	if !ok {
		return vos.Account{}, vos.ErrCPFNotFound
	}
	return account, nil
}
