package account

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type UsecaseMock struct {
	CreateErr    error
	Account      vos.Account
	ListErr      error
	AccountsList []vos.Account
	GetByIDErr   error
	GetByCPFErr error
}

func (u UsecaseMock) Create(ctx context.Context, account vos.CreateInput) (vos.Account, error) {
	return u.Account, u.CreateErr
}

func (u UsecaseMock) GetByID(ctx context.Context, id string) (vos.Account, error) {
	if u.GetByIDErr != nil {
		return vos.Account{}, u.GetByIDErr
	}
	return u.Account,nil
}

func (u UsecaseMock) List() ([]vos.Account, error) {
	return u.AccountsList, u.ListErr
}
