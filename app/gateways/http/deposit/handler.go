package deposit

import "github.com/erikaa81/banco-digital/app/domain/deposit"

type Handler struct {
	UseCase deposit.UseCase
}

func NewHandler(useCase deposit.UseCase) *Handler {
	return &Handler{
		UseCase: useCase,
	}
}
