package vos

import (
	"errors"
	"time"
)

var (
	ErrAccountNotFound            = errors.New("account origin not found")
	ErrAccountOriginNotFound      = errors.New("account origin not found")
	ErrAccountDestinationNotFound = errors.New("account destination not found")
	ErrSameAccounts               = errors.New("the target account must be different from the source account")
	ErrInsufficientBalance        = errors.New("insufficient balance")
)

type Transfer struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
	CreatedAt            time.Time
}

type CreateTransferInput struct {
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
}

type TransferOutput struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               int
	CreatedAt            string
}
