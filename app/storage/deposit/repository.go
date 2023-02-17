package deposit

import "github.com/erikaa81/banco-digital/app/domain/vos"

type Repository struct {
	storage map[string]vos.Deposit
}

func NewRepository() Repository {
	storage := map[string]vos.Deposit{}
	return Repository{storage: storage}
}