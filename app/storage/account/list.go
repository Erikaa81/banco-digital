package account

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r Repository) List(ctx context.Context) ([]vos.Account, error) {
	var accountsList []vos.Account
	for _, account := range r.storage {
		accountsList = append(accountsList, account)
	}
	return accountsList, nil
}
