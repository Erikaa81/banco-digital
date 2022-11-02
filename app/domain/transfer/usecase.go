package transfer

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type UseCase interface {
	Create(context.Context, vos.CreateTransferInput) (vos.Transfer, error)
	List(context.Context, string) ([]vos.Transfer, error)
	GetByID(context.Context, string) (vos.Transfer, error)
}
