package transfer

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r Repository) Store(ctx context.Context, transfer vos.Transfer) (vos.Transfer, error) {
	transfer.ID = uuid.New().String()
	transfer.CreatedAt = time.Now()

	r.storage[transfer.ID] = transfer
	return transfer, nil
}
