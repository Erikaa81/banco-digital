package usecase

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (u Usecase) GetByID(ctx context.Context, id string) (vos.Account, error) {
	if id == "" {
		return vos.Account{}, vos.ErrEmptyID
	}

	account, err := u.repository.GetByID(ctx, id)
	if err != nil {
		return vos.Account{}, err
	}
	return account, nil
}
