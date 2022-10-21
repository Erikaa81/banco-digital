package account

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/erikaa81/banco-digital/app/domain/vos"
	"github.com/erikaa81/banco-digital/app/gateways/http/account/models"
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
		responseError := models.ErrorResponse{Error: "create account request failed: " + err.Error()}
		Respond(w, responseError, http.StatusBadRequest)
		log.Printf("create account request failed %v", err)
		return
	}

	parsingBirthDate, _ := time.Parse(DateLayout, requestBody.BirthDate)

	input := vos.CreateInput{
		Name:      requestBody.Name,
		CPF:       requestBody.CPF,
		Secret:    requestBody.Secret,
		BirthDate: parsingBirthDate,
	}

	output, err := h.UseCase.Create(r.Context(), input)
	if err != nil {
		responseError := models.ErrorResponse{Error: err.Error()}
		Respond(w, responseError, http.StatusBadRequest)
		log.Printf("create account request failed %v", err)
		return
	}

	responseBody := vos.CreateOutput{
		ID:        output.ID,
		Name:      output.Name,
		CPF:       output.CPF,
		CreatedAt: output.CreatedAt.Format(DateLayout),
	}

	Respond(w, responseBody, http.StatusCreated)
}

func Respond(w http.ResponseWriter, responseBody interface{}, statusCode int) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(responseBody)
}
