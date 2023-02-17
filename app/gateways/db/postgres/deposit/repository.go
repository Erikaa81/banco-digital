package deposit

import (
	pgx "github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	*pgx.Pool
}

func NewRepository(db *pgx.Pool) *Repository {
	return &Repository{
		Pool: db,
	}
}