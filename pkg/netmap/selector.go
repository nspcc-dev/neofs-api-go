package netmap

import (
	"fmt"
	"sort"

	"github.com/nspcc-dev/hrw"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

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
	m := c.getSelectionBase(s)
	if len(m) < bucketCount {
		return nil, fmt.Errorf("%w: '%s'", ErrNotEnoughNodes, s.GetName())
	}

	keys := make(sort.StringSlice, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	if len(c.pivot) == 0 {
		// deterministic order in case of zero seed
		keys.Sort()
	}

	nodes := make([]Nodes, 0, len(m))
	for i := range keys {
		ns := m[keys[i]]
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

// getSelectionBase returns nodes grouped by selector attribute.
func (c *Context) getSelectionBase(s *netmap.Selector) map[string]Nodes {
	f := c.Filters[s.GetFilter()]
	isMain := s.GetFilter() == MainFilterName
	result := map[string]Nodes{}
	for i := range c.Netmap.Nodes {
		if isMain || c.match(f, c.Netmap.Nodes[i]) {
			v := c.Netmap.Nodes[i].Attribute(s.GetAttribute())
			result[v] = append(result[v], c.Netmap.Nodes[i])
		}
	}

	if len(c.pivot) != 0 {
		for _, ns := range result {
			hrw.SortSliceByWeightValue(ns, ns.Weights(c.weightFunc), c.pivotHash)
		}
	}
	return result
}
