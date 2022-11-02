package transfer

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r Repository) List(ctx context.Context, accountID string) ([]vos.Transfer, error) {
	var transfersList []vos.Transfer
	for _, transfer := range r.storage {
		if accountID == transfer.AccountOriginID || accountID == transfer.AccountDestinationID {
			transfersList = append(transfersList, transfer)
		}
	}
	return transfersList, nil
}
