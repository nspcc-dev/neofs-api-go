package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXHeader(t *testing.T) {
	x := NewXHeader()

	key := "some key"
	val := "some value"

	x.SetKey(key)
	x.SetValue(val)

	require.Equal(t, key, x.Key())
	require.Equal(t, val, x.Value())

	xV2 := x.ToV2()

	require.Equal(t, key, xV2.GetKey())
	require.Equal(t, val, xV2.GetValue())
}
