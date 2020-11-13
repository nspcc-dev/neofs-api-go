package pkg

import (
	"testing"

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
