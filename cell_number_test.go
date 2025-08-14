package vobj_test

import (
	"testing"

	"github.com/conceitho/vobj"
	"github.com/stretchr/testify/assert"
)

func TestNewCellNumber(t *testing.T) {
	number, err := vobj.NewCellNumber("11992808976")
	assert.Nil(t, err)
	assert.NotNil(t, number)
}

func TestCellNumber_String(t *testing.T) {
	number, err := vobj.NewCellNumber("11992808976")
	assert.Nil(t, err)
	assert.NotNil(t, number)
	assert.Equal(t, "11992808976", number.String())
}

func TestCellNumber_Formatted(t *testing.T) {
	number, err := vobj.NewCellNumber("11992808976")
	assert.Nil(t, err)
	assert.NotNil(t, number)
	assert.Equal(t, "(11) 99280-8976", number.Formatted(), "Número formatado deve ser (11) 99280-8976")
}

func TestNewInvalidCellNumber(t *testing.T) {
	number, err := vobj.NewCellNumber("")
	assert.Empty(t, number)
	assert.NotNil(t, err)
	assert.Equal(t, "número de celular inválido", err.Error())
}
