package usecase

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (u Usecase) List(ctx context.Context) ([]vos.Account, error) {
	accounts, err := u.repository.List(ctx)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
