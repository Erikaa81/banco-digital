package transfer

import "github.com/erikaa81/banco-digital/app/domain/transfer"

type Handler struct {
	UseCase transfer.UseCase
}

func NewHandler(useCase transfer.UseCase) *Handler {
	return &Handler{
		UseCase: useCase,
	}
}
