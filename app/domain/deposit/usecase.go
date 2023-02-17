package deposit

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type UseCase interface {
	Create(context.Context, vos.CreateDepositInput) (vos.Deposit, error)
}
