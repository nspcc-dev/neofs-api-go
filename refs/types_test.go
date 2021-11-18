package refs_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func TestZeroSubnet(t *testing.T) {
	id := new(refs.SubnetID)

	require.True(t, refs.IsZeroSubnet(id))

	id.SetValue(1)
	require.False(t, refs.IsZeroSubnet(id))

	refs.MakeZeroSubnet(id)
	require.True(t, refs.IsZeroSubnet(id))
}
