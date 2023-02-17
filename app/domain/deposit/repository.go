package deposit

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type Repository interface {
	Create(ctx context.Context, deposit vos.Deposit) (vos.Deposit, error)
}
