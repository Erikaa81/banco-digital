package account

import (
	"context"

	pgxType "github.com/jackc/pgx/v4"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r *Repository) GetByID(ctx context.Context, id string) (vos.Account, error) {
	var account vos.Account

	query := "SELECT id, name, cpf, secret, birthdate, created_at FROM accounts WHERE id = $1"

	err := r.Pool.QueryRow(ctx, query, id).Scan(
		&account.ID,
		&account.Name,
		&account.CPF,
		&account.Secret,
		&account.BirthDate,
		&account.CreatedAt,
	)
	if err != nil {
		if err == pgxType.ErrNoRows {
			return vos.Account{}, vos.ErrAccountNotFound
		}
		return vos.Account{}, err
	}
	
	return account, nil
}
