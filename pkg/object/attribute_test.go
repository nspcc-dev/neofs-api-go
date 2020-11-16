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

	require.Equal(t, key, a.Key())
	require.Equal(t, val, a.Value())

	aV2 := a.ToV2()

	require.Equal(t, key, aV2.GetKey())
	require.Equal(t, val, aV2.GetValue())
}

func TestAttributeEncoding(t *testing.T) {
	a := NewAttribute()
	a.SetKey("key")
	a.SetValue("value")

	t.Run("binary", func(t *testing.T) {
		data, err := a.Marshal()
		require.NoError(t, err)

		a2 := NewAttribute()
		require.NoError(t, a2.Unmarshal(data))

		require.Equal(t, a, a2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := a.MarshalJSON()
		require.NoError(t, err)

		a2 := NewAttribute()
		require.NoError(t, a2.UnmarshalJSON(data))

		require.Equal(t, a, a2)
	})
}
