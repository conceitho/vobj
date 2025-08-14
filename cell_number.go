package vobj

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrInvalidCellNumber = errors.New("número de celular inválido")
	// Regex para validar números de celular brasileiros (11 dígitos: DDD + 9)
	phoneRegex     = regexp.MustCompile(`^[1-9]{2}9[0-9]{8}$`)
	nonDigitsRegex = regexp.MustCompile(`\D+`)
)

// CellNumber é um Value Object para números de celular.
type CellNumber struct {
	value string
}

// NewCellNumber cria um novo Value Object CellNumber.
//
// Sanitiza e valida o número, isto é, remove caracteres não numéricos e
// verifica se o número está no formato brasileiro (DDD com 2 dígitos + número com 9 dígitos).
func NewCellNumber(value string) (CellNumber, error) {
	// Remove todos os caracteres não numéricos
	sanitized := nonDigitsRegex.ReplaceAllString(value, "")

	// Valida o formato brasileiro (DDD com 2 dígitos + número com 9 dígitos)
	if !phoneRegex.MatchString(sanitized) {
		return CellNumber{}, ErrInvalidCellNumber
	}

	return CellNumber{value: sanitized}, nil
}

// String retorna o número sanitizado (apenas dígitos).
func (p CellNumber) String() string {
	return p.value
}

// Formatted retorna o número formatado no padrão (XX) XXXXX-XXXX.
func (p CellNumber) Formatted() string {
	if len(p.value) != 11 {
		return p.value
	}
	return fmt.Sprintf("(%s) %s-%s", p.value[0:2], p.value[2:7], p.value[7:11])
}
