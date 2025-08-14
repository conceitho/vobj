package vobj_test

import (
	"testing"

	"github.com/conceitho/vobj"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewPasswordFromPlainText(t *testing.T) {
	pass, err := vobj.NewPasswordFromPlainText("password")
	assert.Nil(t, err, "Erro ao criar senha a partir de texto plano")
	assert.NotNil(t, pass, "Erro ao criar senha a partir de texto plano")
	isEqual := pass.Compare("password")
	assert.True(t, isEqual, "Senha não corresponde ao texto plano")
}

func TestNewPasswordFromHash(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	pass, err := vobj.NewPasswordFromHash(string(hash))
	assert.Nil(t, err)
	assert.NotNil(t, pass, "Erro ao criar senha a partir de hash")
	isEqual := pass.Compare("password")
	assert.True(t, isEqual, "Senha não corresponde ao texto plano")
}

func TestInvalidEmptyPassword(t *testing.T) {
	pass, err := vobj.NewPasswordFromPlainText("")
	assert.NotNil(t, err)
	assert.Empty(t, pass)
	assert.Equal(t, "senha está vazia", err.Error())
}

func TestInvalidEmptyHash(t *testing.T) {
	pass, err := vobj.NewPasswordFromHash("")
	assert.NotNil(t, err)
	assert.Empty(t, pass)
	assert.Equal(t, "hash está vazio", err.Error())
}
