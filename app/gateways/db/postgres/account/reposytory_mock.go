package account

// import (
// 	"context"

// 	"github.com/erikaa81/banco-digital/app/domain/vos"
// )

// type FakeRepository struct {
// 	QueryErr error
// 	Account vos.Account
// }

// func (f FakeRepository) Create(ctx context.Context, account vos.CreateInput) (vos.Account,error) {
// 	return  f.Account, f.QueryErr
// }

// func (f FakeRepository) Scan(dest ...interface{}) error {
// 	return nil
// }



// type PgxIface interface {
// 	Begin(context.Context) (pgx.Tx, error)
// 	Close()
// }
