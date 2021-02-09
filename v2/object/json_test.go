package object_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/stretchr/testify/require"
)

func TestShortHeaderJSON(t *testing.T) {
	h := generateShortHeader("id")

	data, err := h.MarshalJSON()
	require.NoError(t, err)

	h2 := new(object.ShortHeader)
	require.NoError(t, h2.UnmarshalJSON(data))

	require.Equal(t, h, h2)
}

func TestAttributeJSON(t *testing.T) {
	a := generateAttribute("key", "value")

	data, err := a.MarshalJSON()
	require.NoError(t, err)

	a2 := new(object.Attribute)
	require.NoError(t, a2.UnmarshalJSON(data))

	require.Equal(t, a, a2)
}

func TestSplitHeaderJSON(t *testing.T) {
	h := generateSplit("sig")

	data, err := h.MarshalJSON()
	require.NoError(t, err)

	h2 := new(object.SplitHeader)
	require.NoError(t, h2.UnmarshalJSON(data))

	require.Equal(t, h, h2)
}

func TestHeaderJSON(t *testing.T) {
	h := generateHeader(10)

	data, err := h.MarshalJSON()
	require.NoError(t, err)

	h2 := new(object.Header)
	require.NoError(t, h2.UnmarshalJSON(data))

	require.Equal(t, h, h2)
}

func TestHeaderWithSignatureJSON(t *testing.T) {
	h := generateHeaderWithSignature()

	data, err := h.MarshalJSON()
	require.NoError(t, err)

	h2 := new(object.HeaderWithSignature)
	require.NoError(t, h2.UnmarshalJSON(data))

	require.Equal(t, h, h2)
}

func TestObjectJSON(t *testing.T) {
	o := generateObject("data")

	data, err := o.MarshalJSON()
	require.NoError(t, err)

	o2 := new(object.Object)
	require.NoError(t, o2.UnmarshalJSON(data))

	require.Equal(t, o, o2)
}

func TestSearchFilterJSON(t *testing.T) {
	f := generateFilter("key", "value")

	data, err := f.MarshalJSON()
	require.NoError(t, err)

	f2 := new(object.SearchFilter)
	require.NoError(t, f2.UnmarshalJSON(data))

	require.Equal(t, f, f2)
}
