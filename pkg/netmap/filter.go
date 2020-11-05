package netmap

import (
	"fmt"
	"strconv"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// Filter represents v2-compatible netmap filter.
type Filter netmap.Filter

// MainFilterName is a name of the filter
// which points to the whole netmap.
const MainFilterName = "*"

// applyFilter applies named filter to b.
func (c *Context) applyFilter(name string, b *Node) bool {
	return name == MainFilterName || c.match(c.Filters[name], b)
}

// processFilters processes filters and returns error is any of them is invalid.
func (c *Context) processFilters(p *netmap.PlacementPolicy) error {
	for _, f := range p.GetFilters() {
		if err := c.processFilter(f, true); err != nil {
			return err
		}
	}
	return nil
}

func (c *Context) processFilter(f *netmap.Filter, top bool) error {
	if f == nil {
		return fmt.Errorf("%w: FILTER", ErrMissingField)
	}
	if f.GetName() == MainFilterName {
		return fmt.Errorf("%w: '*' is reserved", ErrInvalidFilterName)
	}
	if top && f.GetName() == "" {
		return ErrUnnamedTopFilter
	}
	if !top && f.GetName() != "" && c.Filters[f.GetName()] == nil {
		return fmt.Errorf("%w: '%s'", ErrFilterNotFound, f.GetName())
	}
	switch f.GetOp() {
	case netmap.AND, netmap.OR:
		for _, flt := range f.GetFilters() {
			if err := c.processFilter(flt, false); err != nil {
				return err
			}
		}
	default:
		if len(f.GetFilters()) != 0 {
			return ErrNonEmptyFilters
		} else if !top && f.GetName() != "" { // named reference
			return nil
		}
		switch f.GetOp() {
		case netmap.EQ, netmap.NE:
		case netmap.GT, netmap.GE, netmap.LT, netmap.LE:
			n, err := strconv.ParseUint(f.GetValue(), 10, 64)
			if err != nil {
				return fmt.Errorf("%w: '%s'", ErrInvalidNumber, f.GetValue())
			}
			c.numCache[f] = n
		default:
			return fmt.Errorf("%w: %d", ErrInvalidFilterOp, f.GetOp())
		}
	}
	if top {
		c.Filters[f.GetName()] = f
	}
	return nil
}

// match matches f against b. It returns no errors because
// filter should have been parsed during context creation
// and missing node properties are considered as a regular fail.
func (c *Context) match(f *netmap.Filter, b *Node) bool {
	switch f.GetOp() {
	case netmap.AND, netmap.OR:
		for _, lf := range f.GetFilters() {
			if lf.GetName() != "" {
				lf = c.Filters[lf.GetName()]
			}
			ok := c.match(lf, b)
			if ok == (f.GetOp() == netmap.OR) {
				return ok
			}
		}
		return f.GetOp() == netmap.AND
	default:
		return c.matchKeyValue(f, b)
	}
}

func (c *Context) matchKeyValue(f *netmap.Filter, b *Node) bool {
	switch f.GetOp() {
	case netmap.EQ:
		return b.Attribute(f.GetKey()) == f.GetValue()
	case netmap.NE:
		return b.Attribute(f.GetKey()) != f.GetValue()
	default:
		var attr uint64
		switch f.GetKey() {
		case PriceAttr:
			attr = b.Price
		case CapacityAttr:
			attr = b.Capacity
		default:
			var err error
			attr, err = strconv.ParseUint(b.Attribute(f.GetKey()), 10, 64)
			if err != nil {
				// Note: because filters are somewhat independent from nodes attributes,
				// We don't report an error here, and fail filter instead.
				return false
			}
		}
		switch f.GetOp() {
		case netmap.GT:
			return attr > c.numCache[f]
		case netmap.GE:
			return attr >= c.numCache[f]
		case netmap.LT:
			return attr < c.numCache[f]
		case netmap.LE:
			return attr <= c.numCache[f]
		}
	}
	// will not happen if context was created from f (maybe panic?)
	return false
}

// NewFilter creates and returns new Filter instance.
func NewFilter() *Filter {
	return NewFilterFromV2(new(netmap.Filter))
}

// NewFilterFromV2 converts v2 Filter to Filter.
func NewFilterFromV2(f *netmap.Filter) *Filter {
	return (*Filter)(f)
}

// ToV2 converts Filter to v2 Filter.
func (f *Filter) ToV2() *netmap.Filter {
	return (*netmap.Filter)(f)
}

// Key returns key to filter.
func (f *Filter) Key() string {
	return (*netmap.Filter)(f).
		GetKey()
}

// SetKey sets key to filter.
func (f *Filter) SetKey(key string) {
	(*netmap.Filter)(f).
		SetKey(key)
}

// Value returns value to match.
func (f *Filter) Value() string {
	return (*netmap.Filter)(f).
		GetValue()
}

// SetValue sets value to match.
func (f *Filter) SetValue(val string) {
	(*netmap.Filter)(f).
		SetValue(val)
}

// Name returns filter name.
func (f *Filter) Name() string {
	return (*netmap.Filter)(f).
		GetName()
}

// SetName sets filter name.
func (f *Filter) SetName(name string) {
	(*netmap.Filter)(f).
		SetName(name)
}

// Operation returns filtering operation.
func (f *Filter) Operation() Operation {
	return OperationFromV2(
		(*netmap.Filter)(f).
			GetOp(),
	)
}

// SetOperation sets filtering operation.
func (f *Filter) SetOperation(op Operation) {
	(*netmap.Filter)(f).
		SetOp(op.ToV2())
}

// InnerFilters returns list of inner filters.
func (f *Filter) InnerFilters() []*Filter {
	fs := (*netmap.Filter)(f).
		GetFilters()

	res := make([]*Filter, 0, len(fs))

	for i := range fs {
		res = append(res, NewFilterFromV2(fs[i]))
	}

	return res
}

// SetInnerFilters sets list of inner filters.
func (f *Filter) SetInnerFilters(fs ...*Filter) {
	fsV2 := make([]*netmap.Filter, 0, len(fs))

	for i := range fs {
		fsV2 = append(fsV2, fs[i].ToV2())
	}

	(*netmap.Filter)(f).
		SetFilters(fsV2)
}
