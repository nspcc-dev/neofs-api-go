package netmap

import (
	"errors"

	"github.com/nspcc-dev/hrw"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// Context contains references to named filters and cached numeric values.
type Context struct {
	// Netmap is a netmap structure to operate on.
	Netmap *Netmap
	// Filters stores processed filters.
	Filters map[string]*netmap.Filter
	// Selectors stores processed selectors.
	Selectors map[string]*netmap.Selector
	// Selections stores result of selector processing.
	Selections map[string][]Nodes

	// numCache stores parsed numeric values.
	numCache map[*netmap.Filter]uint64
	// pivot is a seed for HRW.
	pivot []byte
	// pivotHash is a saved HRW hash of pivot
	pivotHash uint64
	// aggregator is returns aggregator determining bucket weight.
	// By default it returns mean value from IQR interval.
	aggregator func() aggregator
	// weightFunc is a weighting function for determining node priority.
	// By default in combines favours low price and high capacity.
	weightFunc weightFunc
}

// Various validation errors.
var (
	ErrMissingField      = errors.New("netmap: nil field")
	ErrInvalidFilterName = errors.New("netmap: filter name is invalid")
	ErrInvalidNumber     = errors.New("netmap: number value expected")
	ErrInvalidFilterOp   = errors.New("netmap: invalid filter operation")
	ErrFilterNotFound    = errors.New("netmap: filter not found")
	ErrNonEmptyFilters   = errors.New("netmap: simple filter must no contain sub-filters")
	ErrNotEnoughNodes    = errors.New("netmap: not enough nodes to SELECT from")
	ErrSelectorNotFound  = errors.New("netmap: selector not found")
	ErrUnnamedTopFilter  = errors.New("netmap: all filters on top level must be named")
)

// NewContext creates new context. It contains various caches.
// In future it may create hierarchical netmap structure to work with.
func NewContext(nm *Netmap) *Context {
	return &Context{
		Netmap:     nm,
		Filters:    make(map[string]*netmap.Filter),
		Selectors:  make(map[string]*netmap.Selector),
		Selections: make(map[string][]Nodes),

		numCache:   make(map[*netmap.Filter]uint64),
		aggregator: newMeanIQRAgg,
		weightFunc: GetDefaultWeightFunc(nm.Nodes),
	}
}

func (c *Context) setPivot(pivot []byte) {
	if len(pivot) != 0 {
		c.pivot = pivot
		c.pivotHash = hrw.Hash(pivot)
	}
}

// GetDefaultWeightFunc returns default weighting function.
func GetDefaultWeightFunc(ns Nodes) weightFunc {
	mean := newMeanAgg()
	min := newMinAgg()
	for i := range ns {
		mean.Add(float64(ns[i].Capacity))
		min.Add(float64(ns[i].Price))
	}
	return newWeightFunc(
		newSigmoidNorm(mean.Compute()),
		newReverseMinNorm(min.Compute()))
}
