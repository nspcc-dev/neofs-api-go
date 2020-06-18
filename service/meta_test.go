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

func TestExtHdrWrapper(t *testing.T) {
	s := wrapExtendedHeaderKV(nil)
	require.Empty(t, s.Key())
	require.Empty(t, s.Value())

	msg := new(RequestExtendedHeader_KV)
	s = wrapExtendedHeaderKV(msg)

	key := "key"
	msg.SetK(key)
	require.Equal(t, key, s.Key())

	val := "val"
	msg.SetV(val)
	require.Equal(t, val, s.Value())
}

func TestRequestExtendedHeader_ExtendedHeaders(t *testing.T) {
	var (
		k1, v1 = "key1", "value1"
		k2, v2 = "key2", "value2"
		h1     = new(RequestExtendedHeader_KV)
		h2     = new(RequestExtendedHeader_KV)
	)

	h1.SetK(k1)
	h1.SetV(v1)

	h2.SetK(k2)
	h2.SetV(v2)

	s := new(RequestExtendedHeader)
	s.SetHeaders([]RequestExtendedHeader_KV{
		*h1, *h2,
	})

	xHdrs := s.ExtendedHeaders()
	require.Len(t, xHdrs, 2)

	require.Equal(t, k1, xHdrs[0].Key())
	require.Equal(t, v1, xHdrs[0].Value())

	require.Equal(t, k2, xHdrs[1].Key())
	require.Equal(t, v2, xHdrs[1].Value())
}
