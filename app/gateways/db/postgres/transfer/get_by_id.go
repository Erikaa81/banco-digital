package transfer

import (
	"context"

	pgxType "github.com/jackc/pgx/v4"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r *Repository) GetByID(ctx context.Context, id string) (vos.Transfer, error) {
	var transfer vos.Transfer

	query := `SELECT id, account_origin_id, account_destination_id, amount, created_at 
	FROM transactions WHERE id = $1`

	err := r.Pool.QueryRow(ctx, query, id).Scan(
		&transfer.ID,
		&transfer.AccountOriginID,
		&transfer.AccountDestinationID,
		&transfer.Amount,
		&transfer.CreatedAt,
	)
	if err != nil {
		if err == pgxType.ErrNoRows {
			return vos.Transfer{}, vos.ErrIDNotFound
		}
		return vos.Transfer{}, err
	}

	return transfer, nil
}
