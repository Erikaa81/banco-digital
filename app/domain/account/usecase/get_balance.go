package usecase

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (u Usecase) GetBalance(ctx context.Context, id string) (int, error) {
	var account vos.Account

	account, err := u.repository.GetByID(ctx, id)
	if err != nil {
		return account.Balance, vos.ErrIDNotFound
	}

	accountBalance, err := u.repository.GetBalance(ctx, id)
	if err != nil {
		return account.Balance, err
	}


	return accountBalance, nil
}
