package transfer

import (
	"context"
	"fmt"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r *Repository) List(ctx context.Context, accountID string) ([]vos.Transfer, error) {
	var transfersList []vos.Transfer
	var transfer vos.Transfer

	rows, err := r.Query(ctx, `SELECT id, account_origin_id, account_destination_id, amount, created_at
	FROM transfers WHERE account_origin_id = $1 OR account_destination_id = $1`, accountID)
	if err != nil {
		return []vos.Transfer{}, err
	}

	for rows.Next() {
		err := rows.Scan(
			&transfer.ID,
			&transfer.AccountOriginID,
			&transfer.AccountDestinationID,
			&transfer.Amount,
			&transfer.CreatedAt,
		)
		if err != nil {
			return []vos.Transfer{}, err
		}
		transfersList = append(transfersList, transfer)
	}
	if len(transfersList) == 0 {
		return []vos.Transfer{}, fmt.Errorf("error while listing transfers: %w", err)
	}

	return transfersList, nil
}
