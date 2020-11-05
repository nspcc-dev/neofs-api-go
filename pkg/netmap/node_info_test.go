package netmap

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

func TestNode_NetworkAddress(t *testing.T) {
	addr := "127.0.0.1:8080"

	nV2 := new(netmap.NodeInfo)
	nV2.SetAddress(addr)

	n := Node{
		InfoV2: nV2,
	}

	require.Equal(t, addr, n.NetworkAddress())
}

func TestNodeStateFromV2(t *testing.T) {
	for _, item := range []struct {
		s   NodeState
		sV2 netmap.NodeState
	}{
		{
			s:   0,
			sV2: netmap.UnspecifiedState,
		},
		{
			s:   NodeStateOnline,
			sV2: netmap.Online,
		},
		{
			s:   NodeStateOffline,
			sV2: netmap.Offline,
		},
	} {
		require.Equal(t, item.s, NodeStateFromV2(item.sV2))
		require.Equal(t, item.sV2, item.s.ToV2())
	}
}

func TestNodeAttributeFromV2(t *testing.T) {
	aV2 := new(netmap.Attribute)
	aV2.SetKey("key")
	aV2.SetValue("value")
	aV2.SetParents([]string{"par1", "par2"})

	a := NewNodeAttributeFromV2(aV2)

	require.Equal(t, aV2, a.ToV2())
}

func TestNodeAttribute_Key(t *testing.T) {
	a := NewNodeAttribute()
	key := "some key"

	a.SetKey(key)

	require.Equal(t, key, a.Key())
}

func TestNodeAttribute_Value(t *testing.T) {
	a := NewNodeAttribute()
	val := "some value"

	a.SetValue(val)

	require.Equal(t, val, a.Value())
}

func TestNodeAttribute_ParentKeys(t *testing.T) {
	a := NewNodeAttribute()
	keys := []string{"par1", "par2"}

	a.SetParentKeys(keys...)

	require.Equal(t, keys, a.ParentKeys())
}
