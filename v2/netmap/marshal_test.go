package netmap_test

import (
	"strconv"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
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

func TestNodeInfo(t *testing.T) {
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

func generateAttribute(k, v string) *netmap.Attribute {
	attr := new(netmap.Attribute)
	attr.SetKey(k)
	attr.SetValue(v)

	return attr
}

func generateNodeInfo(key, addr string, n int) *netmap.NodeInfo {
	nodeInfo := new(netmap.NodeInfo)
	nodeInfo.SetPublicKey([]byte(key))
	nodeInfo.SetAddress(addr)

	attrs := make([]*netmap.Attribute, n)
	for i := 0; i < n; i++ {
		j := strconv.Itoa(n)
		attrs[i] = generateAttribute("key"+j, "value"+j)
	}

	nodeInfo.SetAttributes(attrs)

	return nodeInfo
}
