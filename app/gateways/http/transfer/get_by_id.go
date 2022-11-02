package transfer

import (
	"errors"
	"log"
	"net/http"

	"github.com/erikaa81/banco-digital/app/gateways/http/transfer/models"
	"github.com/gorilla/mux"
)

var (
	ErrIDNotFound = errors.New("id not found")
	ErrSearchID   = errors.New("search by id failed")
)

func (h Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	transferID := mux.Vars(r)["transfer-id"]

	output, err := h.UseCase.GetByID(r.Context(), transferID)
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

	responseBody := models.GetIDTransferResponse{
		ID:                   output.ID,
		AccountOriginID:      output.AccountOriginID,
		AccountDestinationID: output.AccountDestinationID,
		Amount:               output.Amount,
		CreatedAt:            output.CreatedAt.Format(DateLayout),
	}

	Respond(w, responseBody, http.StatusOK)
	log.Printf("get by id request successfully %v", err)
}
