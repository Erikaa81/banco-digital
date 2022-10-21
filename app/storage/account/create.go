package account

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (r Repository) Store(ctx context.Context, account vos.Account) (vos.Account, error) {
	account.ID = uuid.New().String()
	account.CreatedAt = time.Now()

	r.storage[account.ID] = account
	return account, nil
}
