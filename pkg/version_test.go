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

	require.Equal(t, mjr, v.GetMajor())
	require.Equal(t, mnr, v.GetMinor())

	ver := v.ToV2()

	require.Equal(t, mjr, ver.GetMajor())
	require.Equal(t, mnr, ver.GetMinor())
}

func TestSDKVersion(t *testing.T) {
	v := SDKVersion()

	require.Equal(t, uint32(sdkMjr), v.GetMajor())
	require.Equal(t, uint32(sdkMnr), v.GetMinor())
}
