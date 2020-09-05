package netmap

import (
	"fmt"

	"github.com/nspcc-dev/hrw"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// Netmap represents netmap which contains preprocessed nodes.
type Netmap struct {
	Nodes Nodes
}

// NewNetmap constructs netmap from the list of raw nodes.
func NewNetmap(nodes Nodes) (*Netmap, error) {
	return &Netmap{
		Nodes: nodes,
	}, nil
}

func flattenNodes(ns []Nodes) Nodes {
	result := make(Nodes, 0, len(ns))
	for i := range ns {
		result = append(result, ns[i]...)
	}
	return result
}

// GetPlacementVectors returns placement vectors for an object given containerNodes cnt.
func (m *Netmap) GetPlacementVectors(cnt ContainerNodes, pivot []byte) ([]Nodes, error) {
	h := hrw.Hash(pivot)
	wf := GetDefaultWeightFunc(m.Nodes)
	result := make([]Nodes, len(cnt.Replicas()))
	for i, rep := range cnt.Replicas() {
		result[i] = make(Nodes, len(rep))
		copy(result[i], rep)
		hrw.SortSliceByWeightValue(result[i], result[i].Weights(wf), h)
	}
	return result, nil
}

// GetContainerNodes returns nodes corresponding to each replica.
// Order of returned nodes corresponds to order of replicas in p.
// pivot is a seed for HRW sorting.
func (m *Netmap) GetContainerNodes(p *netmap.PlacementPolicy, pivot []byte) (ContainerNodes, error) {
	c := NewContext(m)
	c.setPivot(pivot)
	if err := c.processFilters(p); err != nil {
		return nil, err
	}
	if err := c.processSelectors(p); err != nil {
		return nil, err
	}
	result := make([]Nodes, len(p.GetReplicas()))
	for i, r := range p.GetReplicas() {
		if r == nil {
			return nil, fmt.Errorf("%w: REPLICA", ErrMissingField)
		}
		if r.GetSelector() == "" {
			for _, s := range p.GetSelectors() {
				result[i] = append(result[i], flattenNodes(c.Selections[s.GetName()])...)
			}
		}
		nodes, ok := c.Selections[r.GetSelector()]
		if !ok {
			return nil, fmt.Errorf("%w: REPLICA '%s'", ErrSelectorNotFound, r.GetSelector())
		}
		result[i] = append(result[i], flattenNodes(nodes)...)

	}
	return containerNodes(result), nil
}
