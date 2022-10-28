package account

import (
	"log"
	"net/http"

	"github.com/erikaa81/banco-digital/app/gateways/http/account/models"
)

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	accountsList, err := h.UseCase.List(r.Context())

	var responseError models.ErrorResponse
	if err != nil {
		Respond(w, responseError, http.StatusInternalServerError)
		log.Printf("error when listing accounts")
		return
	}

	accounts := make([]models.Account, 0)
	for _, value := range accountsList {
		account := models.Account{
			ID:        value.ID,
			Name:      value.Name,
			CPF:       value.CPF,
			BirthDate: value.BirthDate,
			CreatedAt: value.CreatedAt,
		}
		accounts = append(accounts, account)
	}

	Respond(w, models.ListResponse{List: accounts}, http.StatusOK)
	log.Printf("list accounts request successfully %v", err)
}
