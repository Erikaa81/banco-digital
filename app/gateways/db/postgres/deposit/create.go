package deposit

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r *Repository) Create(ctx context.Context, deposit vos.Deposit) (vos.Deposit, error) {
	query := `INSERT INTO transactions(
	account_destination_id,
	amount
	) VALUES ($1, $2) returning id, created_at`

	err := r.Pool.QueryRow(ctx, query,
		deposit.AccountDestination_id,
		deposit.Amount,
	).Scan(&deposit.ID, &deposit.CreatedAt)
	if err != nil {
		return vos.Deposit{}, err
	}

	return deposit, nil
}
