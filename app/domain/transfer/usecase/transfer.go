package usecase

import (
	"github.com/erikaa81/banco-digital/app/domain/account"
	"github.com/erikaa81/banco-digital/app/domain/transfer"
)

type Usecase struct {
	repository transfer.Repository
	accountRepository account.Repository
}

func NewTransferUseCase(repository transfer.Repository, accountRepository account.Repository) Usecase {
	return  Usecase{
		repository:        repository,
		accountRepository: accountRepository,
	}
}
