package pkg

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/session"
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

func TestNewXHeaderFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *session.XHeader

		require.Nil(t, NewXHeaderFromV2(x))
	})
}

func TestXHeader_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *XHeader

		require.Nil(t, x.ToV2())
	})
}

func TestNewXHeader(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		xh := NewXHeader()

		// check initial values
		require.Empty(t, xh.Value())
		require.Empty(t, xh.Key())

		// convert to v2 message
		xhV2 := xh.ToV2()

		require.Empty(t, xhV2.GetValue())
		require.Empty(t, xhV2.GetKey())
	})
}
