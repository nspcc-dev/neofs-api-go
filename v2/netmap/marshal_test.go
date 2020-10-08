package netmap_test

import (
	"strconv"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func TestAttribute_StableMarshal(t *testing.T) {
	from := generateAttribute("key", "value")
	transport := new(grpc.NodeInfo_Attribute)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := netmap.AttributeFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestNodeInfo_StableMarshal(t *testing.T) {
	from := generateNodeInfo("publicKey", "/multi/addr", 10)
	transport := new(grpc.NodeInfo)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := netmap.NodeInfoFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestFilter_StableMarshal(t *testing.T) {
	from := generateFilter("key", "value", false)
	transport := new(grpc.Filter)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := netmap.FilterFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestSelector_StableMarshal(t *testing.T) {
	from := generateSelector("name")
	transport := new(grpc.Selector)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := netmap.SelectorFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestReplica_StableMarshal(t *testing.T) {
	from := generateReplica("selector")
	transport := new(grpc.Replica)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := netmap.ReplicaFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestPlacementPolicy_StableMarshal(t *testing.T) {
	from := generatePolicy(3)
	transport := new(grpc.PlacementPolicy)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := netmap.PlacementPolicyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestLocalNodeInfoResponseBody_StableMarshal(t *testing.T) {
	from := generateNodeInfoResponseBody()
	transport := new(grpc.LocalNodeInfoResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := netmap.LocalNodeInfoResponseBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func generateAttribute(k, v string) *netmap.Attribute {
	attr := new(netmap.Attribute)
	attr.SetKey(k)
	attr.SetValue(v)
	attr.SetParents([]string{k, v})

	return attr
}

func generateNodeInfo(key, addr string, n int) *netmap.NodeInfo {
	nodeInfo := new(netmap.NodeInfo)
	nodeInfo.SetPublicKey([]byte(key))
	nodeInfo.SetAddress(addr)
	nodeInfo.SetState(netmap.Online)

	attrs := make([]*netmap.Attribute, n)
	for i := 0; i < n; i++ {
		j := strconv.Itoa(n)
		attrs[i] = generateAttribute("key"+j, "value"+j)
	}

	nodeInfo.SetAttributes(attrs)

	return nodeInfo
}

func generateFilter(key, value string, fin bool) *netmap.Filter {
	f := new(netmap.Filter)
	f.SetKey(key)
	f.SetValue(value)
	f.SetName("name")
	f.SetOp(netmap.AND)
	if !fin {
		ff := generateFilter(key+"fin", value+"fin", true)
		f.SetFilters([]*netmap.Filter{ff})
	} else {
		f.SetFilters([]*netmap.Filter{})
	}

	return f
}

func generateSelector(name string) *netmap.Selector {
	s := new(netmap.Selector)
	s.SetName(name)
	s.SetAttribute("attribute")
	s.SetClause(netmap.Distinct)
	s.SetCount(10)
	s.SetFilter("filter")

	return s
}

func generateReplica(selector string) *netmap.Replica {
	r := new(netmap.Replica)
	r.SetCount(10)
	r.SetSelector(selector)

	return r
}

func generatePolicy(n int) *netmap.PlacementPolicy {
	var (
		p = new(netmap.PlacementPolicy)
		f = make([]*netmap.Filter, 0, n)
		s = make([]*netmap.Selector, 0, n)
		r = make([]*netmap.Replica, 0, n)
	)

	for i := 0; i < n; i++ {
		ind := strconv.Itoa(i)

		f = append(f, generateFilter("key"+ind, "val"+ind, false))
		s = append(s, generateSelector("name"+ind))
		r = append(r, generateReplica("selector"+ind))
	}

	p.SetFilters(f)
	p.SetSelectors(s)
	p.SetReplicas(r)
	p.SetContainerBackupFactor(10)

	return p
}

func generateNodeInfoResponseBody() *netmap.LocalNodeInfoResponseBody {
	ni := generateNodeInfo("key", "/multi/addr", 2)

	r := new(netmap.LocalNodeInfoResponseBody)
	r.SetVersion(generateVersion(2, 1))
	r.SetNodeInfo(ni)

	return r
}

func generateVersion(maj, min uint32) *refs.Version {
	version := new(refs.Version)
	version.SetMajor(maj)
	version.SetMinor(min)

	return version
}
