package transfer

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type RepositoryMock struct {
	TransferOutput    vos.Transfer
	CreateTransferErr error
	Transfer          vos.Transfer
	GetTransferErr    error
	TransferList      []vos.Transfer
	ListErr   error
}

func (r RepositoryMock) Create(context.Context, vos.Transfer) (vos.Transfer, error) {
	return r.TransferOutput, r.CreateTransferErr
}

func (r RepositoryMock) List(context.Context, string) ([]vos.Transfer, error) {
	return r.TransferList, r.ListErr
}

func (r RepositoryMock) GetByID(context.Context, string) (vos.Transfer, error) {
	return r.Transfer, r.GetTransferErr
}
