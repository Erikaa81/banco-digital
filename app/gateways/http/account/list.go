package account

import (
	"errors"
	"log"
	"net/http"

	"github.com/erikaa81/banco-digital/app/gateways/http/account/models"
)

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	accountsList, err := h.UseCase.List()

	var responseError models.ErrorResponse
	if err != nil {
		switch {
		case errors.Is(err, errors.New("error when listing accounts")):
			responseError = models.ErrorResponse{Error: err.Error()}
		default:
			responseError = models.ErrorResponse{Error: err.Error()}
		}

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
	log.Printf("list accounts request failed %v", err)
}
