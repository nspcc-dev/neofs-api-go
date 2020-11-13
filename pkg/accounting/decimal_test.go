package accounting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecimal_Value(t *testing.T) {
	d := NewDecimal()

	v := int64(3)
	d.SetValue(v)

	require.Equal(t, v, d.Value())
}

func TestDecimal_Precision(t *testing.T) {
	d := NewDecimal()

	p := uint32(3)
	d.SetPrecision(p)

	require.Equal(t, p, d.Precision())
}
