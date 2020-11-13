package netmap

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

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

func testNodeAttribute() *NodeAttribute {
	a := new(NodeAttribute)
	a.SetKey("key")
	a.SetValue("value")
	a.SetParentKeys("par1", "par2")

	return a
}

func TestNodeInfoFromV2(t *testing.T) {
	iV2 := new(netmap.NodeInfo)
	iV2.SetPublicKey([]byte{1, 2, 3})
	iV2.SetAddress("456")
	iV2.SetState(netmap.Online)
	iV2.SetAttributes([]*netmap.Attribute{
		testNodeAttribute().ToV2(),
		testNodeAttribute().ToV2(),
	})

	i := NewNodeInfoFromV2(iV2)

	require.Equal(t, iV2, i.ToV2())
}

func TestNodeInfo_PublicKey(t *testing.T) {
	i := new(NodeInfo)
	key := []byte{1, 2, 3}

	i.SetPublicKey(key)

	require.Equal(t, key, i.PublicKey())
}

func TestNodeInfo_Address(t *testing.T) {
	i := new(NodeInfo)
	a := "127.0.0.1:8080"

	i.SetAddress(a)

	require.Equal(t, a, i.Address())
}

func TestNodeInfo_State(t *testing.T) {
	i := new(NodeInfo)
	s := NodeStateOnline

	i.SetState(s)

	require.Equal(t, s, i.State())
}

func TestNodeInfo_Attributes(t *testing.T) {
	i := new(NodeInfo)
	as := []*NodeAttribute{testNodeAttribute(), testNodeAttribute()}

	i.SetAttributes(as...)

	require.Equal(t, as, i.Attributes())
}

func TestNodeAttributeEncoding(t *testing.T) {
	a := testNodeAttribute()

	t.Run("binary", func(t *testing.T) {
		data, err := a.Marshal()
		require.NoError(t, err)

		a2 := NewNodeAttribute()
		require.NoError(t, a2.Unmarshal(data))

		require.Equal(t, a, a2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := a.MarshalJSON()
		require.NoError(t, err)

		a2 := NewNodeAttribute()
		require.NoError(t, a2.UnmarshalJSON(data))

		require.Equal(t, a, a2)
	})
}

func TestNodeInfoEncoding(t *testing.T) {
	i := NewNodeInfo()
	i.SetPublicKey([]byte{1, 2, 3})
	i.SetAddress("192.168.0.1")
	i.SetState(NodeStateOnline)
	i.SetAttributes(testNodeAttribute())

	t.Run("binary", func(t *testing.T) {
		data, err := i.Marshal()
		require.NoError(t, err)

		i2 := NewNodeInfo()
		require.NoError(t, i2.Unmarshal(data))

		require.Equal(t, i, i2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := i.MarshalJSON()
		require.NoError(t, err)

		i2 := NewNodeInfo()
		require.NoError(t, i2.UnmarshalJSON(data))

		require.Equal(t, i, i2)
	})
}
