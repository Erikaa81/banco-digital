package models

type GetBalanceRequest struct {
	ID string `json:"id"`
}

type GetBalanceResponse struct {
	Balance int `json:"balance,omitempty"`
}
