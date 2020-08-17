package netmap

import (
	netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
)

func PlacementPolicyToGRPCMessage(p *PlacementPolicy) *netmap.PlacementPolicy {
	if p == nil {
		return nil
	}

	// TODO: fill me
	return nil
}

func PlacementPolicyFromGRPCMessage(m *netmap.PlacementPolicy) *PlacementPolicy {
	if m == nil {
		return nil
	}

	// TODO: fill me
	return nil
}
