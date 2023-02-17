package account

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type RepositoryMock struct {
	CreateAccountOutput vos.Account
	CreateAccountErr    error
	ListErr             error
	AccountList         []vos.Account
	GetAccountErr       error
	Account             vos.Account
	Balance             int
	GetBalanceErr       error
}

func (r RepositoryMock) GetByCPF(context.Context, string) (vos.Account, error) {
	return r.Account, r.GetAccountErr
}

func (r RepositoryMock) Create(context.Context, vos.Account) (vos.Account, error) {
	return r.CreateAccountOutput, r.CreateAccountErr
}

func (r RepositoryMock) GetByID(context.Context, string) (vos.Account, error) {
	return r.Account, r.GetAccountErr
}

func (r RepositoryMock) List(context.Context) ([]vos.Account, error) {
	return r.AccountList, r.ListErr
}

func (r RepositoryMock) GetBalance(context.Context, string) (int, error) {
	return r.Account.Balance, r.GetBalanceErr
}
