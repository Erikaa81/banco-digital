package models

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrMandatoryName   = errors.New("a name is required to create the account")
	ErrMandatoryCPF    = errors.New("CPF is required to create the account")
	ErrInvalidCPF      = errors.New("CPF is invalid")
	ErrMandatorySecret = errors.New("secret is required to create the account")
	ErrFormat          = errors.New("format is invalid")
)

type CreateRequest struct {
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Secret    string `json:"secret"`
	BirthDate string `json:"birth_date"`
}

func (a *CreateRequest) Validate() error {
	if a.Name == "" {
		return ErrMandatoryName
	}

	if a.CPF == "" {
		return ErrMandatoryCPF
	}

	ValidateCPF := ValidateCPF(a.CPF)
	if !ValidateCPF {
		return ErrInvalidCPF
	}

	if a.Secret == "" {
		return ErrMandatorySecret
	}

	ValidateFormat, ErrFormat := ValidateFormatBirth(a.BirthDate)
	if ValidateFormat {
		return ErrFormat
	}

	return nil
}

type CreateResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	CreatedAt time.Time `json:"created_at"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ValidateCPF(cpf string) bool {
	return len(cpf) == 11
}

func ValidateFormatBirth(birthDate string) (bool, error) {
	_, err := time.Parse("02/01/2006", birthDate)
	if err != nil {
		return true, fmt.Errorf("invalid date format %s", err)
	}
	return false, nil
}
