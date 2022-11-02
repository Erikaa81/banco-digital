package transfer

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r Repository) GetByID(ctx context.Context, id string) (vos.Transfer, error) {
	transfer, ok := r.storage[id]
	if !ok {
		return vos.Transfer{}, vos.ErrIDNotFound
	}
	return transfer, nil
}
