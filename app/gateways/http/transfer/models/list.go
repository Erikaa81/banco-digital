package models

import "time"

type Transfer struct {
	ID                   string    `json:"id"`
	AccountOriginID      string    `json:"account_origin_id"`
	AccountDestinationID string    `json:"account_destination_id"`
	Amount               int       `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

type GetIDResponse struct {
	List []Transfer `json:"transfers"`
}
