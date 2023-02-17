package account

import (
	"context"

	pgxType "github.com/jackc/pgx/v4"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r *Repository) GetBalance(ctx context.Context, id string) (int, error) {
	var balance int

	query := `SELECT SUM(CASE WHEN account_origin_id = $1 then amount * - 1
 				 ELSE amount END) 			 
     FROM  transactions
     where account_destination_id = $1 OR account_origin_id = $1`

	err := r.Pool.QueryRow(ctx, query, id).Scan(
		&balance,
	)
	if err != nil {
		if err == pgxType.ErrNoRows {
			return balance, vos.ErrAccountNotFound
		}
		return balance, err
	}
	return balance, err
}
