package usecase

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (u Usecase) GetByCPF(ctx context.Context, cpf string) (vos.Account, error) {
	if cpf == "" {
		return vos.Account{}, vos.ErrEmptyCPF
	}
	account, err := u.repository.GetByCPF(ctx, cpf)
	if err != nil {
		return vos.Account{}, err
	}
	return account, nil
}
