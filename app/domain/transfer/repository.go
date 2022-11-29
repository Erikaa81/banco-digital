package transfer

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type Repository interface {
	Create(ctx context.Context, transfer vos.Transfer) (vos.Transfer, error)
	List(ctx context.Context, accountID string) ([]vos.Transfer, error)
	GetByID(ctx context.Context, id string) (vos.Transfer, error)
}
