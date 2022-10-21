package account

import "github.com/erikaa81/banco-digital/app/domain/account"

type Handler struct {
	UseCase account.UseCase
}

func NewHandler(useCase account.UseCase) *Handler {
	return &Handler{
		UseCase: useCase,
	}
}
