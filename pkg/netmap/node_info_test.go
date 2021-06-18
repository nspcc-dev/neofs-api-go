package netmap

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	testv2 "github.com/nspcc-dev/neofs-api-go/v2/netmap/test"
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
	t.Run("from nil", func(t *testing.T) {
		var x *netmap.Attribute

		require.Nil(t, NewNodeAttributeFromV2(x))
	})

	t.Run("from non-nil", func(t *testing.T) {
		aV2 := testv2.GenerateAttribute(false)

		a := NewNodeAttributeFromV2(aV2)

		require.Equal(t, aV2, a.ToV2())
	})
}

func TestNodeAttribute_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *NodeAttribute

		require.Nil(t, x.ToV2())
	})
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
	t.Run("from nil", func(t *testing.T) {
		var x *netmap.NodeInfo

		require.Nil(t, NewNodeInfoFromV2(x))
	})

	t.Run("from non-nil", func(t *testing.T) {
		iV2 := testv2.GenerateNodeInfo(false)

		i := NewNodeInfoFromV2(iV2)

		require.Equal(t, iV2, i.ToV2())
	})
}

func TestNodeInfo_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *NodeInfo

		require.Nil(t, x.ToV2())
	})
}

func TestNodeInfo_PublicKey(t *testing.T) {
	i := new(NodeInfo)
	key := []byte{1, 2, 3}

	i.SetPublicKey(key)

	require.Equal(t, key, i.PublicKey())
}

func TestNodeInfo_IterateAddresses(t *testing.T) {
	i := new(NodeInfo)

	as := []string{"127.0.0.1:8080", "127.0.0.1:8081"}

	i.SetAddresses(as...)

	as2 := make([]string, 0, i.NumberOfAddresses())

	IterateAllAddresses(i, func(addr string) {
		as2 = append(as2, addr)
	})

	require.Equal(t, as, as2)
	require.EqualValues(t, len(as), i.NumberOfAddresses())
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
	i.SetAddresses("192.168.0.1", "192.168.0.2")
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

func TestNewNodeAttribute(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		attr := NewNodeAttribute()

		// check initial values
		require.Empty(t, attr.Key())
		require.Empty(t, attr.Value())
		require.Nil(t, attr.ParentKeys())

		// convert to v2 message
		attrV2 := attr.ToV2()

		require.Empty(t, attrV2.GetKey())
		require.Empty(t, attrV2.GetValue())
		require.Nil(t, attrV2.GetParents())
	})
}

func TestNewNodeInfo(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		ni := NewNodeInfo()

		// check initial values
		require.Nil(t, ni.PublicKey())

		require.Zero(t, ni.NumberOfAddresses())
		require.Nil(t, ni.Attributes())
		require.Zero(t, ni.State())

		// convert to v2 message
		niV2 := ni.ToV2()

		require.Nil(t, niV2.GetPublicKey())
		require.Zero(t, niV2.NumberOfAddresses())
		require.Nil(t, niV2.GetAttributes())
		require.EqualValues(t, netmap.UnspecifiedState, niV2.GetState())
	})
}

func TestNodeState_String(t *testing.T) {
	toPtr := func(v NodeState) *NodeState {
		return &v
	}

	testEnumStrings(t, new(NodeState), []enumStringItem{
		{val: toPtr(NodeStateOnline), str: "ONLINE"},
		{val: toPtr(NodeStateOffline), str: "OFFLINE"},
		{val: toPtr(0), str: "UNSPECIFIED"},
	})
}
