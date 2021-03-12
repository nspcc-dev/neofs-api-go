package netmaptest

import (
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	refstest "github.com/nspcc-dev/neofs-api-go/v2/refs/test"
	sessiontest "github.com/nspcc-dev/neofs-api-go/v2/session/test"
)

func GenerateFilter(empty bool) *netmap.Filter {
	return generateFilter(empty, true)
}

func generateFilter(empty, withSub bool) *netmap.Filter {
	m := new(netmap.Filter)

	if !empty {
		m.SetKey("filter key")
		m.SetValue("filter value")
		m.SetName("filter name")
		m.SetOp(1)

		if withSub {
			m.SetFilters([]*netmap.Filter{
				generateFilter(empty, false),
				generateFilter(empty, false),
			})
		}
	}

	return m
}

func GenerateFilters(empty bool) (res []*netmap.Filter) {
	if !empty {
		res = append(res,
			GenerateFilter(false),
			GenerateFilter(false),
		)
	}

	return
}

func GenerateSelector(empty bool) *netmap.Selector {
	m := new(netmap.Selector)

	if !empty {
		m.SetCount(66)
		m.SetAttribute("selector attribute")
		m.SetFilter("select filter")
		m.SetName("select name")
		m.SetClause(1)
	}

	return m
}

func GenerateSelectors(empty bool) (res []*netmap.Selector) {
	if !empty {
		res = append(res,
			GenerateSelector(false),
			GenerateSelector(false),
		)
	}

	return
}

func GenerateReplica(empty bool) *netmap.Replica {
	m := new(netmap.Replica)

	if !empty {
		m.SetCount(42)
		m.SetSelector("replica selector")
	}

	return m
}

func GenerateReplicas(empty bool) (res []*netmap.Replica) {
	if !empty {
		res = append(res,
			GenerateReplica(false),
			GenerateReplica(false),
		)
	}

	return
}

func GeneratePlacementPolicy(empty bool) *netmap.PlacementPolicy {
	m := new(netmap.PlacementPolicy)

	if !empty {
		m.SetContainerBackupFactor(322)
	}

	m.SetFilters(GenerateFilters(empty))
	m.SetSelectors(GenerateSelectors(empty))
	m.SetReplicas(GenerateReplicas(empty))

	return m
}

func GenerateAttribute(empty bool) *netmap.Attribute {
	m := new(netmap.Attribute)

	if !empty {
		m.SetKey("attribute key")
		m.SetValue("attribute val")
	}

	return m
}

func GenerateAttributes(empty bool) (res []*netmap.Attribute) {
	if !empty {
		res = append(res,
			GenerateAttribute(false),
			GenerateAttribute(false),
		)
	}

	return
}

func GenerateNodeInfo(empty bool) *netmap.NodeInfo {
	m := new(netmap.NodeInfo)

	if !empty {
		m.SetAddress("node address")
		m.SetPublicKey([]byte{1, 2, 3})
		m.SetState(33)
	}

	m.SetAttributes(GenerateAttributes(empty))

	return m
}

func GenerateLocalNodeInfoRequestBody(empty bool) *netmap.LocalNodeInfoRequestBody {
	m := new(netmap.LocalNodeInfoRequestBody)

	return m
}

func GenerateLocalNodeInfoRequest(empty bool) *netmap.LocalNodeInfoRequest {
	m := new(netmap.LocalNodeInfoRequest)

	m.SetBody(GenerateLocalNodeInfoRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateLocalNodeInfoResponseBody(empty bool) *netmap.LocalNodeInfoResponseBody {
	m := new(netmap.LocalNodeInfoResponseBody)

	m.SetVersion(refstest.GenerateVersion(empty))
	m.SetNodeInfo(GenerateNodeInfo(empty))

	return m
}

func GenerateLocalNodeInfoResponse(empty bool) *netmap.LocalNodeInfoResponse {
	m := new(netmap.LocalNodeInfoResponse)

	m.SetBody(GenerateLocalNodeInfoResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateNetworkInfo(empty bool) *netmap.NetworkInfo {
	m := new(netmap.NetworkInfo)

	if !empty {
		m.SetMagicNumber(228)
		m.SetCurrentEpoch(666)
	}

	return m
}

func GenerateNetworkInfoRequestBody(empty bool) *netmap.NetworkInfoRequestBody {
	m := new(netmap.NetworkInfoRequestBody)

	return m
}

func GenerateNetworkInfoRequest(empty bool) *netmap.NetworkInfoRequest {
	m := new(netmap.NetworkInfoRequest)

	m.SetBody(GenerateNetworkInfoRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateNetworkInfoResponseBody(empty bool) *netmap.NetworkInfoResponseBody {
	m := new(netmap.NetworkInfoResponseBody)

	m.SetNetworkInfo(GenerateNetworkInfo(empty))

	return m
}

func GenerateNetworkInfoResponse(empty bool) *netmap.NetworkInfoResponse {
	m := new(netmap.NetworkInfoResponse)

	m.SetBody(GenerateNetworkInfoResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}
