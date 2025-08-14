package vobj_test

import (
	"testing"
	"time"

	"github.com/conceitho/vobj"
	"github.com/stretchr/testify/assert"
)

func TestNewPeriod(t *testing.T) {
	expected := time.Now().Format("012006")
	formated := time.Now().Format("01/2006")
	period, err := vobj.NewPeriod(time.Now())
	assert.Nil(t, err)
	assert.NotNil(t, period)
	assert.Equal(t, expected, period.String())
	assert.Equal(t, formated, period.Formatted())
}
