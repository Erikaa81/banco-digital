package transfer

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r *Repository) Create(ctx context.Context, transfer vos.Transfer) (vos.Transfer, error) {
	query := `INSERT INTO transactions(
	account_origin_id,
	account_destination_id,
	amount
	) VALUES ($1, $2, $3) returning id, created_at`

	err := r.Pool.QueryRow(ctx, query,
		transfer.AccountOriginID,
		transfer.AccountDestinationID,
		transfer.Amount,
	).Scan(&transfer.ID, &transfer.CreatedAt)
	if err != nil {
		return vos.Transfer{}, err
	}
	
	return transfer, nil
}
