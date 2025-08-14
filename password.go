package vobj

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// Password é um Value Object que gerencia o hash de senhas.
type Password struct {
	hash string
}

// NewPasswordFromPlainText cria um novo Value Object Password a partir de uma senha em texto plano.
// Ele gera o hash automaticamente.
func NewPasswordFromPlainText(plainText string) (Password, error) {
	if plainText == "" {
		return Password{}, errors.New("senha está vazia")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, err
	}
	return Password{hash: string(hash)}, nil
}

// NewPasswordFromHash cria um Value Object Password a partir de um hash já existente (do banco de dados).
// Não há validação aqui, pois confiamos que o hash no banco é válido.
func NewPasswordFromHash(hash string) (Password, error) {
	if hash == "" {
		return Password{}, errors.New("hash está vazio")
	}
	return Password{hash: hash}, nil
}

// Compare compara a senha em texto plano com o hash armazenado.
func (p Password) Compare(plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(plainText))
	return err == nil
}

// String retorna o valor do hash.
func (p Password) String() string {
	return p.hash
}
