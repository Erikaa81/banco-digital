package account

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type Repository interface {
	GetByCPF(ctx context.Context, cpf string) (vos.Account, error)
	Create(ctx context.Context, account vos.Account) (vos.Account, error)
	GetByID(ctx context.Context, id string) (vos.Account, error)
	List(ctx context.Context) ([]vos.Account, error)
}
