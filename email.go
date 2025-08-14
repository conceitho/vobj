package vobj

import (
	"errors"
	"regexp"
	"strings"
)

var (
	// ErrInvalidEmail é retornado quando o formato do email é inválido.
	ErrInvalidEmail = errors.New("formato de email inválido")
	emailRegex      = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

type Email struct {
	value string
}

// NewEmail é o construtor para o Value Object Email. Ele valida o formato.
func NewEmail(value string) (Email, error) {
	if !emailRegex.MatchString(strings.ToLower(value)) {
		return Email{}, ErrInvalidEmail
	}
	return Email{value: value}, nil
}

// String retorna o valor primitivo do email.
func (e Email) String() string {
	return e.value
}
