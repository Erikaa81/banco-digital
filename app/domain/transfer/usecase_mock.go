package transfer

import (
	"context"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

type UsecaseMock struct {
	CreateErr          error
	Transfer           vos.Transfer
	GetByIDErr         error
	TransferList       []vos.Transfer
	GetByIDTransferErr error
}

func (u UsecaseMock) Create(ctx context.Context, account vos.CreateTransferInput) (vos.Transfer, error) {
	return u.Transfer, u.CreateErr
}

func (u UsecaseMock) List(ctx context.Context, accountID string) ([]vos.Transfer, error) {
	return u.TransferList, u.GetByIDErr
}

func (u UsecaseMock) GetByID(ctx context.Context, transferID string) (vos.Transfer, error) {
	if u.GetByIDTransferErr != nil {
		return vos.Transfer{}, u.GetByIDTransferErr
	}
	return u.Transfer, nil
}
