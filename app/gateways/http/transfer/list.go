package transfer

import (
	"log"
	"net/http"

	"github.com/erikaa81/banco-digital/app/gateways/http/transfer/models"
	"github.com/gorilla/mux"
)

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	accountID := mux.Vars(r)["account-id"]

	transfersList, err := h.UseCase.List(r.Context(), accountID)

	var responseError models.ErrorResponse
	if err != nil {
		Respond(w, responseError, http.StatusInternalServerError)
		log.Printf("error when listing transfers")
		return
	}

	transfers := make([]models.Transfer, 0)
	for _, value := range transfersList {
		transfer := models.Transfer{
			ID:                   value.ID,
			AccountOriginID:      value.AccountOriginID,
			AccountDestinationID: value.AccountDestinationID,
			Amount:               value.Amount,
			CreatedAt:            value.CreatedAt,
		}
		transfers = append(transfers, transfer)
	}

	Respond(w, models.GetIDResponse{List: transfers}, http.StatusOK)
	log.Printf("list transfers request successfully")
}
