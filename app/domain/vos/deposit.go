package vos

import (
	"time"
)

type Deposit struct {
	ID                    string
	AccountDestination_id string
	Amount                int
	CreatedAt             time.Time
}

type CreateDepositInput struct {
	AccountDestination_id string
	Amount                int
}

type DepositOutput struct {
	ID                    string
	AccountDestination_id string
	Amount                int
	CreatedAt             string
}
