package transfer

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type Repository interface {
	Store(context.Context, vos.Transfer) (vos.Transfer, error)
	List(context.Context, string) ([]vos.Transfer, error)
	GetByID(context.Context, string) (vos.Transfer, error)
}
