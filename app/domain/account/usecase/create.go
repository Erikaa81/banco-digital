package usecase

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

func (u Usecase) Create(ctx context.Context, input vos.CreateInput) (vos.Account, error) {
	persistedAccount, err := u.repository.GetByCPF(ctx, input.CPF)
	if err != nil && !errors.Is(err, vos.ErrAccountNotFound) {
		return vos.Account{}, err
	}

	if !reflect.DeepEqual(persistedAccount, vos.Account{}) {
		return vos.Account{}, vos.ErrCPFAlreadyExists
	}

	if err = validateAge(input.BirthDate); err != nil {
		return vos.Account{}, err
	}

	account := vos.Account{
		Name:      input.Name,
		CPF:       input.CPF,
		Secret:    input.Secret,
		BirthDate: input.BirthDate,
	}

	output, err := u.repository.Create(ctx, account)
	if err != nil {
		return vos.Account{}, err
	}
	return output, nil
}

var ErrInvalidAge = errors.New("the client must be of legal age")

func validateAge(birthDate time.Time) error {
	now := time.Now()
	ageOf := time.Date(now.Year()-18, now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())

	if birthDate.After(ageOf) {
		return ErrInvalidAge
	}
	return nil
}
