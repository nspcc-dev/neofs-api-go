package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// PlacementPolicy represents v2-compatible placement policy.
type PlacementPolicy netmap.PlacementPolicy

// PlacementPolicyToJSON encodes PlacementPolicy to JSON format.
func PlacementPolicyToJSON(p *PlacementPolicy) ([]byte, error) {
	return netmap.PlacementPolicyToJSON(p.ToV2())
}

// PlacementPolicyFromJSON decodes PlacementPolicy from JSON format.
func PlacementPolicyFromJSON(data []byte) (*PlacementPolicy, error) {
	p, err := netmap.PlacementPolicyFromJSON(data)
	if err != nil {
		return nil, err
	}

	return NewPlacementPolicyFromV2(p), nil
}

// NewPlacementPolicy creates and returns new PlacementPolicy instance.
func NewPlacementPolicy() *PlacementPolicy {
	return NewPlacementPolicyFromV2(new(netmap.PlacementPolicy))
}

// NewPlacementPolicyFromV2 converts v2 PlacementPolicy to PlacementPolicy.
func NewPlacementPolicyFromV2(f *netmap.PlacementPolicy) *PlacementPolicy {
	return (*PlacementPolicy)(f)
}

// ToV2 converts PlacementPolicy to v2 PlacementPolicy.
func (p *PlacementPolicy) ToV2() *netmap.PlacementPolicy {
	return (*netmap.PlacementPolicy)(p)
}

// Replicas returns list of object replica descriptors.
func (p *PlacementPolicy) Replicas() []*Replica {
	rs := (*netmap.PlacementPolicy)(p).
		GetReplicas()

	res := make([]*Replica, 0, len(rs))

	for i := range rs {
		res = append(res, NewReplicaFromV2(rs[i]))
	}

	return res
}

// SetReplicas sets list of object replica descriptors.
func (p *PlacementPolicy) SetReplicas(rs ...*Replica) {
	rsV2 := make([]*netmap.Replica, 0, len(rs))

	for i := range rs {
		rsV2 = append(rsV2, rs[i].ToV2())
	}

	(*netmap.PlacementPolicy)(p).
		SetReplicas(rsV2)
}

// ContainerBackupFactor returns container backup factor.
func (p *PlacementPolicy) ContainerBackupFactor() uint32 {
	return (*netmap.PlacementPolicy)(p).
		GetContainerBackupFactor()
}

// SetContainerBackupFactor sets container backup factor.
func (p *PlacementPolicy) SetContainerBackupFactor(f uint32) {
	(*netmap.PlacementPolicy)(p).
		SetContainerBackupFactor(f)
}

// Selector returns set of selectors to form the container's nodes subset.
func (p *PlacementPolicy) Selectors() []*Selector {
	rs := (*netmap.PlacementPolicy)(p).
		GetSelectors()

	res := make([]*Selector, 0, len(rs))

	for i := range rs {
		res = append(res, NewSelectorFromV2(rs[i]))
	}

	return res
}

// SetSelectors sets set of selectors to form the container's nodes subset.
func (p *PlacementPolicy) SetSelectors(ss ...*Selector) {
	rsV2 := make([]*netmap.Selector, 0, len(ss))

	for i := range ss {
		rsV2 = append(rsV2, ss[i].ToV2())
	}

	(*netmap.PlacementPolicy)(p).
		SetSelectors(rsV2)
}

// Filters returns list of named filters to reference in selectors.
func (p *PlacementPolicy) Filters() []*Filter {
	return filtersFromV2(
		(*netmap.PlacementPolicy)(p).
			GetFilters(),
	)
}

// SetFilters sets list of named filters to reference in selectors.
func (p *PlacementPolicy) SetFilters(fs ...*Filter) {
	(*netmap.PlacementPolicy)(p).
		SetFilters(filtersToV2(fs))
}
