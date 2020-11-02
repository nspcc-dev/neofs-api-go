package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// fixme: make types instead of aliases to v2 structures
type PlacementPolicy = netmap.PlacementPolicy
type Selector = netmap.Selector
type Filter = netmap.Filter
type Replica = netmap.Replica
type Clause = netmap.Clause
type Operation = netmap.Operation

func PlacementPolicyToJSON(p *PlacementPolicy) ([]byte, error) {
	return netmap.PlacementPolicyToJSON(p)
}

func PlacementPolicyFromJSON(data []byte) (*PlacementPolicy, error) {
	return netmap.PlacementPolicyFromJSON(data)
}
