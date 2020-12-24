package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewVersionFromV2(t *testing.T) {
	v := NewVersion()

	var mjr, mnr uint32 = 1, 2

	v.SetMajor(mjr)
	v.SetMinor(mnr)

	require.Equal(t, mjr, v.Major())
	require.Equal(t, mnr, v.Minor())

	ver := v.ToV2()

	require.Equal(t, mjr, ver.GetMajor())
	require.Equal(t, mnr, ver.GetMinor())
}

func TestSDKVersion(t *testing.T) {
	v := SDKVersion()

	require.Equal(t, uint32(sdkMjr), v.Major())
	require.Equal(t, uint32(sdkMnr), v.Minor())
}

func TestIsSupportedVersion(t *testing.T) {
	require.Error(t, IsSupportedVersion(nil))

	v := NewVersion()

	v.SetMajor(1)
	require.Error(t, IsSupportedVersion(v))

	v.SetMajor(3)
	require.Error(t, IsSupportedVersion(v))

	for _, item := range []struct {
		mjr, maxMnr uint32
	}{
		{
			mjr:    2,
			maxMnr: 1,
		},
	} {
		v.SetMajor(item.mjr)

		for i := uint32(0); i < item.maxMnr; i++ {
			v.SetMinor(i)

			require.NoError(t, IsSupportedVersion(v))
		}

		v.SetMinor(item.maxMnr + 1)

		require.Error(t, IsSupportedVersion(v))
	}
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
