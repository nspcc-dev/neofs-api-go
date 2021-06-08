package pkg

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func TestSignatureEncoding(t *testing.T) {
	s := NewSignature()
	s.SetKey([]byte("key"))
	s.SetSign([]byte("sign"))

	t.Run("binary", func(t *testing.T) {
		data, err := s.Marshal()
		require.NoError(t, err)

		s2 := NewSignature()
		require.NoError(t, s2.Unmarshal(data))

		require.Equal(t, s, s2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := s.MarshalJSON()
		require.NoError(t, err)

		s2 := NewSignature()
		require.NoError(t, s2.UnmarshalJSON(data))

		require.Equal(t, s, s2)
	})
}

func TestNewSignatureFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *refs.Signature

		require.Nil(t, NewSignatureFromV2(x))
	})
}

func TestSignature_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *Signature

		require.Nil(t, x.ToV2())
	})
}

func TestNewSignature(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		sg := NewSignature()

		// check initial values
		require.Nil(t, sg.Key())
		require.Nil(t, sg.Sign())

		// convert to v2 message
		sgV2 := sg.ToV2()

		require.Nil(t, sgV2.GetKey())
		require.Nil(t, sgV2.GetSign())
	})
}
