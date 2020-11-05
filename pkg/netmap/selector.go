package netmap

import (
	"fmt"
	"sort"

	"github.com/nspcc-dev/hrw"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// Selector represents v2-compatible netmap selector.
type Selector netmap.Selector

// processSelectors processes selectors and returns error is any of them is invalid.
func (c *Context) processSelectors(p *netmap.PlacementPolicy) error {
	for _, s := range p.GetSelectors() {
		if s == nil {
			return fmt.Errorf("%w: SELECT", ErrMissingField)
		} else if s.GetFilter() != MainFilterName {
			_, ok := c.Filters[s.GetFilter()]
			if !ok {
				return fmt.Errorf("%w: SELECT FROM '%s'", ErrFilterNotFound, s.GetFilter())
			}
		}
		c.Selectors[s.GetName()] = s
		result, err := c.getSelection(p, s)
		if err != nil {
			return err
		}
		c.Selections[s.GetName()] = result
	}
	return nil
}

// GetNodesCount returns amount of buckets and nodes in every bucket
// for a given placement policy.
func GetNodesCount(p *netmap.PlacementPolicy, s *netmap.Selector) (int, int) {
	switch s.GetClause() {
	case netmap.Same:
		return 1, int(p.GetContainerBackupFactor() * s.GetCount())
	default:
		return int(s.GetCount()), int(p.GetContainerBackupFactor())
	}
}

// getSelection returns nodes grouped by s.attribute.
func (c *Context) getSelection(p *netmap.PlacementPolicy, s *netmap.Selector) ([]Nodes, error) {
	bucketCount, nodesInBucket := GetNodesCount(p, s)
	buckets := c.getSelectionBase(s)
	if len(buckets) < bucketCount {
		return nil, fmt.Errorf("%w: '%s'", ErrNotEnoughNodes, s.GetName())
	}

	if len(c.pivot) == 0 {
		// Deterministic order in case of zero seed.
		if s.GetAttribute() == "" {
			sort.Slice(buckets, func(i, j int) bool {
				return buckets[i].nodes[0].ID < buckets[j].nodes[0].ID
			})
		} else {
			sort.Slice(buckets, func(i, j int) bool {
				return buckets[i].attr < buckets[j].attr
			})
		}
	}

	nodes := make([]Nodes, 0, len(buckets))
	for i := range buckets {
		ns := buckets[i].nodes
		if len(ns) >= nodesInBucket {
			nodes = append(nodes, ns[:nodesInBucket])
		}
	}
	if len(nodes) < bucketCount {
		return nil, fmt.Errorf("%w: '%s'", ErrNotEnoughNodes, s.GetName())
	}
	if len(c.pivot) != 0 {
		weights := make([]float64, len(nodes))
		for i := range nodes {
			weights[i] = GetBucketWeight(nodes[i], c.aggregator(), c.weightFunc)
		}
		hrw.SortSliceByWeightIndex(nodes, weights, c.pivotHash)
	}
	return nodes[:bucketCount], nil
}

type nodeAttrPair struct {
	attr  string
	nodes Nodes
}

// getSelectionBase returns nodes grouped by selector attribute.
// It it guaranteed that each pair will contain at least one node.
func (c *Context) getSelectionBase(s *netmap.Selector) []nodeAttrPair {
	f := c.Filters[s.GetFilter()]
	isMain := s.GetFilter() == MainFilterName
	result := []nodeAttrPair{}
	nodeMap := map[string]Nodes{}
	attr := s.GetAttribute()
	for i := range c.Netmap.Nodes {
		if isMain || c.match(f, c.Netmap.Nodes[i]) {
			if attr == "" {
				// Default attribute is transparent identifier which is different for every node.
				result = append(result, nodeAttrPair{attr: "", nodes: Nodes{c.Netmap.Nodes[i]}})
			} else {
				v := c.Netmap.Nodes[i].Attribute(attr)
				nodeMap[v] = append(nodeMap[v], c.Netmap.Nodes[i])
			}
		}
	}
	if attr != "" {
		for k, ns := range nodeMap {
			result = append(result, nodeAttrPair{attr: k, nodes: ns})
		}
	}

	if len(c.pivot) != 0 {
		for i := range result {
			hrw.SortSliceByWeightValue(result[i].nodes, result[i].nodes.Weights(c.weightFunc), c.pivotHash)
		}
	}
	return result
}

// NewSelector creates and returns new Selector instance.
func NewSelector() *Selector {
	return NewSelectorFromV2(new(netmap.Selector))
}

// NewSelectorFromV2 converts v2 Selector to Selector.
func NewSelectorFromV2(f *netmap.Selector) *Selector {
	return (*Selector)(f)
}

// ToV2 converts Selector to v2 Selector.
func (s *Selector) ToV2() *netmap.Selector {
	return (*netmap.Selector)(s)
}

// Name returns selector name.
func (s *Selector) Name() string {
	return (*netmap.Selector)(s).
		GetName()
}

// SetName sets selector name.
func (s *Selector) SetName(name string) {
	(*netmap.Selector)(s).
		SetName(name)
}

// Count returns count of nodes to select from bucket.
func (s *Selector) Count() uint32 {
	return (*netmap.Selector)(s).
		GetCount()
}

// SetCount sets count of nodes to select from bucket.
func (s *Selector) SetCount(c uint32) {
	(*netmap.Selector)(s).
		SetCount(c)
}

// Clause returns modifier showing how to form a bucket.
func (s *Selector) Clause() Clause {
	return ClauseFromV2(
		(*netmap.Selector)(s).
			GetClause(),
	)
}

// SetClause sets modifier showing how to form a bucket.
func (s *Selector) SetClause(c Clause) {
	(*netmap.Selector)(s).
		SetClause(c.ToV2())
}

// Attribute returns attribute bucket to select from.
func (s *Selector) Attribute() string {
	return (*netmap.Selector)(s).
		GetAttribute()
}

// SetAttribute sets attribute bucket to select from.
func (s *Selector) SetAttribute(a string) {
	(*netmap.Selector)(s).
		SetAttribute(a)
}

// Filter returns filter reference to select from.
func (s *Selector) Filter() string {
	return (*netmap.Selector)(s).
		GetFilter()
}

// SetFilter sets filter reference to select from.
func (s *Selector) SetFilter(f string) {
	(*netmap.Selector)(s).
		SetFilter(f)
}
