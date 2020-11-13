package netmap_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

func TestFilterJSON(t *testing.T) {
	f := generateFilter("key", "value", false)

	d, err := f.MarshalJSON()
	require.NoError(t, err)

	f2 := new(netmap.Filter)
	require.NoError(t, f2.UnmarshalJSON(d))

	require.Equal(t, f, f2)
}

func TestSelectorJSON(t *testing.T) {
	s := generateSelector("name")

	data, err := s.MarshalJSON()
	require.NoError(t, err)

	s2 := new(netmap.Selector)
	require.NoError(t, s2.UnmarshalJSON(data))

	require.Equal(t, s, s2)
}

func TestReplicaJSON(t *testing.T) {
	s := generateReplica("selector")

	data, err := s.MarshalJSON()
	require.NoError(t, err)

	s2 := new(netmap.Replica)
	require.NoError(t, s2.UnmarshalJSON(data))

	require.Equal(t, s, s2)
}

func TestAttributeJSON(t *testing.T) {
	a := generateAttribute("key", "value")

	data, err := a.MarshalJSON()
	require.NoError(t, err)

	a2 := new(netmap.Attribute)
	require.NoError(t, a2.UnmarshalJSON(data))

	require.Equal(t, a, a2)
}

func TestNodeInfoJSON(t *testing.T) {
	i := generateNodeInfo("key", "value", 3)

	data, err := i.MarshalJSON()
	require.NoError(t, err)

	i2 := new(netmap.NodeInfo)
	require.NoError(t, i2.UnmarshalJSON(data))

	require.Equal(t, i, i2)
}
