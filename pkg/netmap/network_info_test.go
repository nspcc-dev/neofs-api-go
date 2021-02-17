package netmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNetworkInfo_CurrentEpoch(t *testing.T) {
	i := NewNetworkInfo()
	e := uint64(13)

	i.SetCurrentEpoch(e)

	require.Equal(t, e, i.CurrentEpoch())
	require.Equal(t, e, i.ToV2().GetCurrentEpoch())
}

func TestNetworkInfo_MagicNumber(t *testing.T) {
	i := NewNetworkInfo()
	m := uint64(666)

	i.SetMagicNumber(m)

	require.Equal(t, m, i.MagicNumber())
	require.Equal(t, m, i.ToV2().GetMagicNumber())
}

func TestNetworkInfoEncoding(t *testing.T) {
	i := NewNetworkInfo()
	i.SetCurrentEpoch(13)
	i.SetMagicNumber(666)

	t.Run("binary", func(t *testing.T) {
		data, err := i.Marshal()
		require.NoError(t, err)

		i2 := NewNetworkInfo()
		require.NoError(t, i2.Unmarshal(data))

		require.Equal(t, i, i2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := i.MarshalJSON()
		require.NoError(t, err)

		i2 := NewNetworkInfo()
		require.NoError(t, i2.UnmarshalJSON(data))

		require.Equal(t, i, i2)
	})
}
