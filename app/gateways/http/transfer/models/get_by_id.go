package models

type GetIDTransferRequest struct {
	ID string `json:"id"`
}

type GetIDTransferResponse struct {
	ID                   string `json:"id"`
	AccountOriginID      string `json:"account_origin_id"`
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
	CreatedAt            string `json:"created_at"`
}
