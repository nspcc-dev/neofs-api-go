package netmap_test

import (
	"testing"

	netmaptest "github.com/nspcc-dev/neofs-api-go/v2/netmap/test"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/message"
	messagetest "github.com/nspcc-dev/neofs-api-go/v2/rpc/message/test"
)

func TestMessageConvert(t *testing.T) {
	messagetest.TestRPCMessage(t,
		func(empty bool) message.Message { return netmaptest.GenerateFilter(empty) },
		func(empty bool) message.Message { return netmaptest.GenerateSelector(empty) },
		func(empty bool) message.Message { return netmaptest.GenerateReplica(empty) },
		func(empty bool) message.Message { return netmaptest.GeneratePlacementPolicy(empty) },
		func(empty bool) message.Message { return netmaptest.GenerateAttribute(empty) },
		func(empty bool) message.Message { return netmaptest.GenerateNodeInfo(empty) },
		func(empty bool) message.Message { return netmaptest.GenerateLocalNodeInfoRequest(empty) },
		func(empty bool) message.Message { return netmaptest.GenerateLocalNodeInfoResponseBody(empty) },
		func(empty bool) message.Message { return netmaptest.GenerateNetworkParameter(empty) },
		func(empty bool) message.Message { return netmaptest.GenerateNetworkConfig(empty) },
		func(empty bool) message.Message { return netmaptest.GenerateNetworkInfo(empty) },
		func(empty bool) message.Message { return netmaptest.GenerateNetworkInfoRequest(empty) },
		func(empty bool) message.Message { return netmaptest.GenerateNetworkInfoResponseBody(empty) },
	)
}
