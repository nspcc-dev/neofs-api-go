package netmap

import (
	"errors"

	"github.com/nspcc-dev/hrw"
)

// Context contains references to named filters and cached numeric values.
type Context struct {
	// Netmap is a netmap structure to operate on.
	Netmap *Netmap
	// Filters stores processed filters.
	Filters map[string]*Filter
	// Selectors stores processed selectors.
	Selectors map[string]*Selector
	// Selections stores result of selector processing.
	Selections map[string][]Nodes

	// numCache stores parsed numeric values.
	numCache map[*Filter]uint64
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
	// container backup factor is a factor for selector counters that expand
	// amount of chosen nodes.
	cbf uint32
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
		Filters:    make(map[string]*Filter),
		Selectors:  make(map[string]*Selector),
		Selections: make(map[string][]Nodes),

		numCache:   make(map[*Filter]uint64),
		aggregator: newMeanIQRAgg,
		weightFunc: GetDefaultWeightFunc(nm.Nodes),
		cbf:        defaultCBF,
	}
}

func (c *Context) setPivot(pivot []byte) {
	if len(pivot) != 0 {
		c.pivot = pivot
		c.pivotHash = hrw.Hash(pivot)
	}
}

func (c *Context) setCBF(cbf uint32) {
	if cbf == 0 {
		c.cbf = defaultCBF
	} else {
		c.cbf = cbf
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
