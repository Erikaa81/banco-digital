package models

import "time"

type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
}
type ListResponse struct {
	List []Account `json:"accounts"`
}
