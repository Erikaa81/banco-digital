package account

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type Repository interface {
	GetByCPF(context.Context, string) (vos.Account, error)
	Store(context.Context, vos.Account) (vos.Account, error)
	GetByID(context.Context, string) (vos.Account, error)
	List() ([]vos.Account, error)
}
