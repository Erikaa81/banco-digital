package models

type GetIDRequest struct {
	ID string `json:"id"`
}

type GetIDResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	CreatedAt string `json:"created_at"`
}
