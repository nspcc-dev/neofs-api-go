package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequestMetaHeader_SetEpoch(t *testing.T) {
	m := new(ResponseMetaHeader)
	epoch := uint64(3)
	m.SetEpoch(epoch)
	require.Equal(t, epoch, m.GetEpoch())
}

func TestRequestMetaHeader_SetVersion(t *testing.T) {
	m := new(ResponseMetaHeader)
	version := uint32(3)
	m.SetVersion(version)
	require.Equal(t, version, m.GetVersion())
}

func TestRequestMetaHeader_SetRaw(t *testing.T) {
	m := new(RequestMetaHeader)

	m.SetRaw(true)
	require.True(t, m.GetRaw())

	m.SetRaw(false)
	require.False(t, m.GetRaw())
}
