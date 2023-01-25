package account

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r *Repository) Create(ctx context.Context, account vos.Account) (vos.Account, error) {
	query := `INSERT INTO accounts(
	name,
	cpf,
	secret,
	birthdate
	) VALUES ($1, $2, $3, $4) returning id, created_at`

	err := r.QueryRow(context.Background(), query,
		account.Name,
		account.CPF,
		account.Secret,
		account.BirthDate,
	).Scan(&account.ID, &account.CreatedAt)
	if err != nil {
		return vos.Account{}, err
	}

	return account, nil
}



// func (r *Repository) Create(db PgxIface, account vos.Account) (vos.Account, error) {
// 	tx, err := db.Begin(context.Background())

// 	query := `INSERT INTO accounts(
// 	name,
// 	cpf,
// 	secret,
// 	birthdate
// 	) VALUES ($1, $2, $3, $4) returning id, created_at`

// 	if err = tx.QueryRow(context.Background(), query,
// 		account.Name,
// 		account.CPF,
// 		account.Secret,
// 		account.BirthDate,
// 	).Scan(&account.ID, &account.CreatedAt); err != nil {
// 		return vos.Account{}, err
// 	}

// 	return account, nil
// }
