package object

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAttribute(t *testing.T) {
	key, val := "some key", "some value"

	a := NewAttribute()
	a.SetKey(key)
	a.SetValue(val)

	require.Equal(t, key, a.GetKey())
	require.Equal(t, val, a.GetValue())

	aV2 := a.ToV2()

	require.Equal(t, key, aV2.GetKey())
	require.Equal(t, val, aV2.GetValue())
}
