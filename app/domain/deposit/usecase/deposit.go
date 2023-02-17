package usecase

import (
	"github.com/erikaa81/banco-digital/app/domain/account"
	"github.com/erikaa81/banco-digital/app/domain/deposit"
)

type Usecase struct {
	repository        deposit.Repository
	accountRepository account.Repository
}

func NewDepositUseCase(repository deposit.Repository, accountRepository account.Repository) Usecase {
	return Usecase{
		repository:        repository,
		accountRepository: accountRepository,
	}
}
