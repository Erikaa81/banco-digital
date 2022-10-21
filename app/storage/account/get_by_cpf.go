package account

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r Repository) GetByCPF(ctx context.Context, cpf string) (vos.Account, error) {
	for _, account := range r.storage {
		if account.CPF == cpf {
			return account, nil
		}
	}
	return vos.Account{}, vos.ErrCPFNotFound
}
