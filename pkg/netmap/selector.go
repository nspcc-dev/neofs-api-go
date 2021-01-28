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
func (c *Context) processSelectors(p *PlacementPolicy) error {
	for _, s := range p.Selectors() {
		if s == nil {
			return fmt.Errorf("%w: SELECT", ErrMissingField)
		} else if s.Filter() != MainFilterName {
			_, ok := c.Filters[s.Filter()]
			if !ok {
				return fmt.Errorf("%w: SELECT FROM '%s'", ErrFilterNotFound, s.Filter())
			}
		}

		c.Selectors[s.Name()] = s

		result, err := c.getSelection(p, s)
		if err != nil {
			return err
		}

		c.Selections[s.Name()] = result
	}

	return nil
}

// GetNodesCount returns amount of buckets and minimum number of nodes in every bucket
// for the given selector.
func GetNodesCount(_ *PlacementPolicy, s *Selector) (int, int) {
	switch s.Clause() {
	case ClauseSame:
		return 1, int(s.Count())
	default:
		return int(s.Count()), 1
	}
}

// getSelection returns nodes grouped by s.attribute.
// Last argument specifies if more buckets can be used to fullfill CBF.
func (c *Context) getSelection(p *PlacementPolicy, s *Selector) ([]Nodes, error) {
	bucketCount, nodesInBucket := GetNodesCount(p, s)
	buckets := c.getSelectionBase(s)

	if len(buckets) < bucketCount {
		return nil, fmt.Errorf("%w: '%s'", ErrNotEnoughNodes, s.Name())
	}

	if len(c.pivot) == 0 {
		// Deterministic order in case of zero seed.
		if s.Attribute() == "" {
			sort.Slice(buckets, func(i, j int) bool {
				return buckets[i].nodes[0].ID < buckets[j].nodes[0].ID
			})
		} else {
			sort.Slice(buckets, func(i, j int) bool {
				return buckets[i].attr < buckets[j].attr
			})
		}
	}

	maxNodesInBucket := nodesInBucket * int(p.ContainerBackupFactor())
	nodes := make([]Nodes, 0, len(buckets))
	fallback := make([]Nodes, 0, len(buckets))

	for i := range buckets {
		ns := buckets[i].nodes
		if len(ns) >= maxNodesInBucket {
			nodes = append(nodes, ns[:maxNodesInBucket])
		} else if len(ns) >= nodesInBucket {
			fallback = append(fallback, ns)
		}
	}

	if len(nodes) < bucketCount {
		// Fallback to using minimum allowed backup factor (1).
		nodes = append(nodes, fallback...)
		if len(nodes) < bucketCount {
			return nil, fmt.Errorf("%w: '%s'", ErrNotEnoughNodes, s.Name())
		}
	}

	if len(c.pivot) != 0 {
		weights := make([]float64, len(nodes))
		for i := range nodes {
			weights[i] = GetBucketWeight(nodes[i], c.aggregator(), c.weightFunc)
		}

		hrw.SortSliceByWeightIndex(nodes, weights, c.pivotHash)
	}

	if s.Attribute() == "" {
		nodes, fallback = nodes[:bucketCount], nodes[bucketCount:]
		for i := range fallback {
			index := i % bucketCount
			if len(nodes[index]) >= maxNodesInBucket {
				break
			}
			nodes[index] = append(nodes[index], fallback[i]...)
		}
	}

	return nodes[:bucketCount], nil
}

type nodeAttrPair struct {
	attr  string
	nodes Nodes
}

// getSelectionBase returns nodes grouped by selector attribute.
// It it guaranteed that each pair will contain at least one node.
func (c *Context) getSelectionBase(s *Selector) []nodeAttrPair {
	f := c.Filters[s.Filter()]
	isMain := s.Filter() == MainFilterName
	result := []nodeAttrPair{}
	nodeMap := map[string]Nodes{}
	attr := s.Attribute()

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

// Marshal marshals Selector into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (s *Selector) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*netmap.Selector)(s).StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Selector.
func (s *Selector) Unmarshal(data []byte) error {
	return (*netmap.Selector)(s).
		Unmarshal(data)
}

// MarshalJSON encodes Selector to protobuf JSON format.
func (s *Selector) MarshalJSON() ([]byte, error) {
	return (*netmap.Selector)(s).
		MarshalJSON()
}

// UnmarshalJSON decodes Selector from protobuf JSON format.
func (s *Selector) UnmarshalJSON(data []byte) error {
	return (*netmap.Selector)(s).
		UnmarshalJSON(data)
}
