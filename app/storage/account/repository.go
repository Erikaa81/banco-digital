package account

import "github.com/erikaa81/banco-digital/app/domain/vos"

type Repository struct {
	storage map[string]vos.Account
}

func NewRepository() Repository {
	storage := map[string]vos.Account{}
	return Repository{storage: storage}
}
