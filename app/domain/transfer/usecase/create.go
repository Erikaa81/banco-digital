package usecase

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (u Usecase) Create(ctx context.Context, input vos.CreateTransferInput) (vos.Transfer, error) {
	accountOrigin, err := u.accountRepository.GetByID(ctx, input.AccountOriginID)
	if err != nil {
		return vos.Transfer{}, vos.ErrAccountOriginNotFound
	}

	accountDestination, err := u.accountRepository.GetByID(ctx, input.AccountDestinationID)
	if err != nil {
		return vos.Transfer{}, vos.ErrAccountDestinationNotFound
	}

	accountOrigin.Balance, err = u.accountRepository.GetBalance(ctx, input.AccountOriginID)
	if err != nil {
		return vos.Transfer{}, vos.ErrInsufficientBalance
	}

	transfer := vos.Transfer{
		AccountOriginID:      accountOrigin.ID,
		AccountDestinationID: accountDestination.ID,
		Amount:               input.Amount,
	}

	if &accountOrigin.Balance == nil || accountOrigin.Balance < transfer.Amount {
		return vos.Transfer{}, vos.ErrInsufficientBalance
	}

	output, err := u.repository.Create(ctx, transfer)
	if err != nil {
		return vos.Transfer{}, err
	}

	return output, nil
}
