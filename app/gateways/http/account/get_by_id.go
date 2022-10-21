package account

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/erikaa81/banco-digital/app/gateways/http/account/models"
)

var (
	ErrIDNotFound = errors.New("id not found")
	ErrSearchID = errors.New("search by id failed")
)

func (h Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	output, err := h.UseCase.GetByID(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, ErrIDNotFound):
			responseError := models.ErrorResponse{Error: "id not found"}
			Respond(w, responseError, http.StatusNotFound)
			log.Printf("id not found %v", err)

		default:
			responseError := models.ErrorResponse{Error: "search by id failed"}
			Respond(w, responseError, http.StatusInternalServerError)
			log.Printf("search by id failed %v", err)
		}
		return
	}

	responseBody := models.GetIDResponse{
		ID:        output.ID,
		Name:      output.Name,
		CPF:       output.CPF,
		CreatedAt: output.CreatedAt.Format(DateLayout),
	}

	Respond(w, responseBody, http.StatusOK)
}
