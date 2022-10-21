package usecase

import (
	"github.com/erikaa81/banco-digital/app/domain/account"
)

type Usecase struct {
	repository account.Repository
}

func NewAccounteUseCase(repository  account.Repository) Usecase {
	return Usecase{repository: repository}
}
