package transfer

import "github.com/erikaa81/banco-digital/app/domain/vos"

type Repository struct {
	storage map[string]vos.Transfer
}

func NewRepository() Repository {
	storage := map[string]vos.Transfer{}
	return Repository{storage: storage}
}