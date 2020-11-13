package container_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/container"
	"github.com/stretchr/testify/require"
)

func TestContainerJSON(t *testing.T) {
	c := generateContainer("nonce")

	data, err := c.MarshalJSON()
	require.NoError(t, err)

	c2 := new(container.Container)
	require.NoError(t, c2.UnmarshalJSON(data))

	require.Equal(t, c, c2)
}

func TestAttributeJSON(t *testing.T) {
	b := generateAttribute("key", "value")

	data, err := b.MarshalJSON()
	require.NoError(t, err)

	b2 := new(container.Attribute)
	require.NoError(t, b2.UnmarshalJSON(data))

	require.Equal(t, b, b2)
}
