package vobj_test

import (
	"testing"

	"github.com/conceitho/vobj"
	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	expected := "teste@teste.com"
	email, err := vobj.NewEmail(expected)
	assert.Nil(t, err)
	assert.NotNil(t, email)
	assert.Equal(t, expected, email.String(), "Email deve ser igual ao valor fornecido")
}

func TestInvalidateEmail(t *testing.T) {
	email, err := vobj.NewEmail("teste@teste")
	assert.NotNil(t, err)
	assert.Empty(t, email)
}
