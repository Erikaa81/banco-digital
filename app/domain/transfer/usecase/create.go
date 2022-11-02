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

	transfer := vos.Transfer{
		AccountOriginID:      accountOrigin.ID,
		AccountDestinationID: accountDestination.ID,
		Amount:               input.Amount,
	}

	output, err := u.repository.Store(ctx, transfer)
	if err != nil {
		return vos.Transfer{}, err
	}
	return output, nil
}
