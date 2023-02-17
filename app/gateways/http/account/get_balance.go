package account

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/erikaa81/banco-digital/app/gateways/http/account/models"
)


func (h Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	output, err := h.UseCase.GetBalance(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, ErrIDNotFound):
			responseError := models.ErrorResponse{Error: "id not found"}
			Respond(w, responseError, http.StatusNotFound)
			log.Printf("id not found %v", err)

		default:
			responseError := models.ErrorResponse{Error: "search balance failed"}
			Respond(w, responseError, http.StatusInternalServerError)
			log.Printf("search balance failed %v", err)
		}
		return
	}

	responseBody := models.GetBalanceResponse{
		Balance:   output,
	}

	Respond(w, responseBody, http.StatusOK)
	log.Printf("get balance request successfully")
}
