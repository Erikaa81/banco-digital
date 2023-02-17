package usecase

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (u Usecase) Create(ctx context.Context, input vos.CreateDepositInput) (vos.Deposit, error) {

	accountDestination, err := u.accountRepository.GetByID(ctx, input.AccountDestination_id)
	if err != nil {
		return vos.Deposit{}, vos.ErrAccountDestinationNotFound
	}

	deposit := vos.Deposit{
		AccountDestination_id: accountDestination.ID,
		Amount:                input.Amount,
	}

	output, err := u.repository.Create(ctx, deposit)
	if err != nil {
		return vos.Deposit{}, err
	}

	return output, nil
}
