package vos

import (
	"errors"
	"time"
)

var (
	ErrCPFNotFound      = errors.New("cpf not found")
	ErrCPFAlreadyExists = errors.New("there is already an account for this cpf")
	ErrIDNotFound       = errors.New("id not found")
	ErrEmptyID          = errors.New("the id was not filled")
	ErrEmptyCPF          = errors.New("the cpf was not filled")
)

type Account struct {
	ID        string
	Name      string
	CPF       string
	Secret    string
	BirthDate time.Time
	CreatedAt time.Time
}

type CreateInput struct {
	Name      string
	CPF       string
	Secret    string
	BirthDate time.Time
}

type CreateOutput struct {
	ID        string
	Name      string
	CPF       string
	CreatedAt string
}
