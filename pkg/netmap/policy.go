package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// fixme: make types instead of aliases to v2 structures
type PlacementPolicy = netmap.PlacementPolicy
type Replica = netmap.Replica

func PlacementPolicyToJSON(p *PlacementPolicy) ([]byte, error) {
	return netmap.PlacementPolicyToJSON(p)
}

func PlacementPolicyFromJSON(data []byte) (*PlacementPolicy, error) {
	return netmap.PlacementPolicyFromJSON(data)
}
