package usecase

import "github.com/erikaa81/banco-digital/app/domain/vos"

func (u Usecase) List() ([]vos.Account, error) {
	accounts, err := u.repository.List()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
