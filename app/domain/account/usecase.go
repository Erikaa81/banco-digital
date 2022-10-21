package account

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type UseCase interface {
	Create(context.Context, vos.CreateInput) (vos.Account, error)
	GetByID(context.Context, string) (vos.Account, error)
	List() ([]vos.Account, error)
}
