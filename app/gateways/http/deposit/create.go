package deposit

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/erikaa81/banco-digital/app/domain/vos"
	"github.com/erikaa81/banco-digital/app/gateways/http/deposit/models"
)

var DateLayout = "02/01/2006"

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var requestBody models.CreateRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		responseError := models.ErrorResponse{Error: "error when decoding body"}
		Respond(w, responseError, http.StatusInternalServerError)
		log.Printf("error when decoding body")
		return
	}

	err = requestBody.Validate()
	if err != nil {
		responseError := models.ErrorResponse{Error: "create deposit request failed: " + err.Error()}
		Respond(w, responseError, http.StatusBadRequest)
		log.Printf("create deposit request failed %v", err)
		return
	}

	input := vos.CreateDepositInput{
		AccountDestination_id: requestBody.AccountDestinationID,
		Amount:                requestBody.Amount,
	}

	output, err := h.UseCase.Create(r.Context(), input)
	if err != nil {
		responseError := models.ErrorResponse{Error: err.Error()}
		Respond(w, responseError, http.StatusBadRequest)
		log.Printf("create deposit request failed %v", err)
		return
	}

	responseBody := vos.DepositOutput{
		ID:                    output.ID,
		AccountDestination_id: output.AccountDestination_id,
		Amount:                output.Amount,
		CreatedAt:             output.CreatedAt.Format(DateLayout),
	}

	Respond(w, responseBody, http.StatusCreated)
	log.Printf("create deposit request successfully")
}

func Respond(w http.ResponseWriter, responseBody interface{}, statusCode int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(responseBody)
}
