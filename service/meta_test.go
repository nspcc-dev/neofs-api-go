package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCutRestoreMeta(t *testing.T) {
	items := []func() SeizedMetaHeaderContainer{
		func() SeizedMetaHeaderContainer {
			m := new(RequestMetaHeader)
			m.SetEpoch(1)
			return m
		},
	}

	for _, item := range items {
		v1 := item()
		m1 := v1.CutMeta()
		v1.RestoreMeta(m1)

		require.Equal(t, item(), v1)
	}
}

func TestRequestExtendedHeader_KV_Setters(t *testing.T) {
	s := new(RequestExtendedHeader_KV)

	key := "key"
	s.SetK(key)
	require.Equal(t, key, s.GetK())

	val := "val"
	s.SetV(val)
	require.Equal(t, val, s.GetV())
}

func TestRequestExtendedHeader_SetHeaders(t *testing.T) {
	s := new(RequestExtendedHeader)

	hdr := RequestExtendedHeader_KV{}
	hdr.SetK("key")
	hdr.SetV("val")

	hdrs := []RequestExtendedHeader_KV{
		hdr,
	}

	s.SetHeaders(hdrs)

	require.Equal(t, hdrs, s.GetHeaders())
}
