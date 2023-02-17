package models

import (
	"errors"
	"time"
)

type CreateRequest struct {
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
}

var (
	ErrMandatoryAccountDestination = errors.New("account destination is required to create the transfer")
	ErrAmountInvalid               = errors.New("amount invalid")
)

func (t *CreateRequest) Validate() error {
	if t.AccountDestinationID == "" {
		return ErrMandatoryAccountDestination
	}

	if t.Amount <= 0 {
		return ErrAmountInvalid
	}
	return nil
}

type CreateResponse struct {
	ID                   string    `json:"id"`
	AccountDestinationID string    `json:"account_destination_id"`
	Amount               int       `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
