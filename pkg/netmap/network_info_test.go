package netmap_test

import (
	"testing"

	. "github.com/nspcc-dev/neofs-api-go/pkg/netmap"
	netmaptest "github.com/nspcc-dev/neofs-api-go/pkg/netmap/test"
	"github.com/stretchr/testify/require"
)

func TestNetworkParameter_Key(t *testing.T) {
	i := NewNetworkParameter()

	k := []byte("key")

	i.SetKey(k)

	require.Equal(t, k, i.Key())
	require.Equal(t, k, i.ToV2().GetKey())
}

func TestNetworkParameter_Value(t *testing.T) {
	i := NewNetworkParameter()

	v := []byte("value")

	i.SetValue(v)

	require.Equal(t, v, i.Value())
	require.Equal(t, v, i.ToV2().GetValue())
}

func TestNewNetworkParameterFromV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		require.Nil(t, NewNetworkParameterFromV2(nil))
	})
}

func TestNetworkParameter_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *NetworkParameter

		require.Nil(t, x.ToV2())
	})
}

func TestNewNetworkParameter(t *testing.T) {
	x := NewNetworkParameter()

	// check initial values
	require.Nil(t, x.Key())
	require.Nil(t, x.Value())

	// convert to v2 message
	xV2 := x.ToV2()

	require.Nil(t, xV2.GetKey())
	require.Nil(t, xV2.GetValue())
}

func TestNetworkConfig_SetParameters(t *testing.T) {
	x := NewNetworkConfig()

	require.Zero(t, x.NumberOfParameters())

	called := 0

	x.IterateParameters(func(p *NetworkParameter) bool {
		called++
		return false
	})

	require.Zero(t, called)

	pps := []*NetworkParameter{
		netmaptest.NetworkParameter(),
		netmaptest.NetworkParameter(),
	}

	x.SetParameters(pps...)

	require.EqualValues(t, len(pps), x.NumberOfParameters())

	var dst []*NetworkParameter

	x.IterateParameters(func(p *NetworkParameter) bool {
		dst = append(dst, p)
		called++
		return false
	})

	require.Equal(t, pps, dst)
	require.Equal(t, len(pps), called)
}

func TestNewNetworkConfigFromV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		require.Nil(t, NewNetworkConfigFromV2(nil))
	})
}

func TestNetworkConfig_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {

		var x *NetworkConfig
		require.Nil(t, x.ToV2())
	})
}

func TestNewNetworkConfig(t *testing.T) {
	x := NewNetworkConfig()

	// check initial values
	require.Zero(t, x.NumberOfParameters())

	// convert to v2 message
	xV2 := x.ToV2()

	require.Zero(t, xV2.NumberOfParameters())
}

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

func TestNetworkInfo_MsPerBlock(t *testing.T) {
	i := NewNetworkInfo()

	const ms = 987

	i.SetMsPerBlock(ms)

	require.EqualValues(t, ms, i.MsPerBlock())
	require.EqualValues(t, ms, i.ToV2().GetMsPerBlock())
}

func TestNetworkInfo_Config(t *testing.T) {
	i := NewNetworkInfo()

	c := netmaptest.NetworkConfig()

	i.SetNetworkConfig(c)

	require.Equal(t, c, i.NetworkConfig())
}

func TestNetworkInfoEncoding(t *testing.T) {
	i := netmaptest.NetworkInfo()

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

func TestNewNetworkInfoFromV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		require.Nil(t, NewNetworkInfoFromV2(nil))
	})
}

func TestNetworkInfo_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *NetworkInfo

		require.Nil(t, x.ToV2())
	})
}

func TestNewNetworkInfo(t *testing.T) {
	ni := NewNetworkInfo()

	// check initial values
	require.Zero(t, ni.CurrentEpoch())
	require.Zero(t, ni.MagicNumber())
	require.Zero(t, ni.MsPerBlock())

	// convert to v2 message
	niV2 := ni.ToV2()

	require.Zero(t, niV2.GetCurrentEpoch())
	require.Zero(t, niV2.GetMagicNumber())
	require.Zero(t, niV2.GetMsPerBlock())
}
