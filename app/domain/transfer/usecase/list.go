package usecase

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (u Usecase) List(ctx context.Context, accountID string) ([]vos.Transfer, error) {
	transfers, err := u.repository.List(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}
