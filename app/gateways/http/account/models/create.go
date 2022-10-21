package models

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrMandatoryName      = errors.New("a name is required to create the account")
	ErrMandatoryCPF       = errors.New("CPF is required to create the account")
	ErrInvalidCPF         = errors.New("CPF is invalid")
	ErrMandatorySecret    = errors.New("secret is required to create the account")
	ErrInvalidDateFormat  = errors.New("format date is invalid")
	ErrMandatoryBirthDate = errors.New("birth date is required to create the account")
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

	ValidateCPF := validateCPF(a.CPF)
	if !ValidateCPF {
		return ErrInvalidCPF
	}

	if a.Secret == "" {
		return ErrMandatorySecret
	}

	err := validateFormatBirth(a.BirthDate)
	if err != nil {
		return ErrInvalidDateFormat
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

func validateCPF(cpf string) bool {
	return len(cpf) == 11
}

func validateFormatBirth(birthDate string) error {
	if strings.TrimSpace(birthDate) == "" {
		return ErrInvalidDateFormat
	}

	_, err := time.Parse("02/01/2006", birthDate)
	if err != nil {
		return fmt.Errorf("invalid date format %s", err)
	}
	return nil
}
