package deposit

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r Repository) Store(ctx context.Context, deposit vos.Deposit) (vos.Deposit, error) {
	deposit.ID = uuid.New().String()
	deposit.CreatedAt = time.Now()

	r.storage[deposit.ID] = deposit
	return deposit, nil
}