package pkg

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func TestNewVersion(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		v := NewVersion()

		// check initial values
		require.Zero(t, v.Major())
		require.Zero(t, v.Minor())

		// convert to v2 message
		vV2 := v.ToV2()

		require.Empty(t, vV2.GetMajor())
		require.Empty(t, vV2.GetMinor())
	})

	t.Run("setting values", func(t *testing.T) {
		v := NewVersion()

		var mjr, mnr uint32 = 1, 2

		v.SetMajor(mjr)
		v.SetMinor(mnr)

		require.Equal(t, mjr, v.Major())
		require.Equal(t, mnr, v.Minor())

		ver := v.ToV2()

		require.Equal(t, mjr, ver.GetMajor())
		require.Equal(t, mnr, ver.GetMinor())
	})
}

func TestSDKVersion(t *testing.T) {
	v := SDKVersion()

	require.Equal(t, uint32(sdkMjr), v.Major())
	require.Equal(t, uint32(sdkMnr), v.Minor())
}

func TestVersionEncoding(t *testing.T) {
	v := NewVersion()
	v.SetMajor(1)
	v.SetMinor(2)

	t.Run("binary", func(t *testing.T) {
		data, err := v.Marshal()
		require.NoError(t, err)

		v2 := NewVersion()
		require.NoError(t, v2.Unmarshal(data))

		require.Equal(t, v, v2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := v.MarshalJSON()
		require.NoError(t, err)

		v2 := NewVersion()
		require.NoError(t, v2.UnmarshalJSON(data))

		require.Equal(t, v, v2)
	})
}

func TestNewVersionFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *refs.Version

		require.Nil(t, NewVersionFromV2(x))
	})
}

func TestVersion_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *Version

		require.Nil(t, x.ToV2())
	})
}
