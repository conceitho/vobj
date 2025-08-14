package vobj

import (
	"fmt"
	"time"
)

// Period representa uma competência no formato MMYYYY.
type Period struct {
	value string
}

func NewPeriod(value time.Time) (Period, error) {
	return Period{
		value: value.Format("012006"),
	}, nil
}

func (p *Period) String() string {
	return p.value
}

// Formatted retorna a competência formatada no padrão MM/YYYY.
func (p Period) Formatted() string {
	if len(p.value) != 6 {
		return p.value
	}
	return fmt.Sprintf("%s/%s", p.value[0:2], p.value[2:6])
}
