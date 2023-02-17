package account

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type UsecaseMock struct {
	CreateErr     error
	Account       vos.Account
	ListErr       error
	AccountsList  []vos.Account
	GetByIDErr    error
	GetByCPFErr   error
	Balance       int
	GetBalanceErr error
}

func (u UsecaseMock) Create(ctx context.Context, account vos.CreateInput) (vos.Account, error) {
	return u.Account, u.CreateErr
}

func (u UsecaseMock) GetByID(ctx context.Context, id string) (vos.Account, error) {
	if u.GetByIDErr != nil {
		return vos.Account{}, u.GetByIDErr
	}
	return u.Account, nil
}

func (u UsecaseMock) List(ctx context.Context) ([]vos.Account, error) {
	return u.AccountsList, u.ListErr
}

func (u UsecaseMock) GetBalance(ctx context.Context, id string) (int, error) {
	if u.GetBalanceErr != nil {
		return u.Account.Balance, u.GetBalanceErr
	}
	return u.Account.Balance, nil
}
