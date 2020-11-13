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
