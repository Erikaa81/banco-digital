package usecase

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (u Usecase) GetByID(ctx context.Context, id string) (vos.Transfer, error) {
	if id == "" {
		return vos.Transfer{}, vos.ErrEmptyID
	}

	transfer, err := u.repository.GetByID(ctx, id)
	if err != nil {
		return vos.Transfer{}, err
	}
	return transfer, nil
}
//conferir as transaçoes da conta