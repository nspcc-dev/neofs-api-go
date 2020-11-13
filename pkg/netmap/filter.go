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
func (c *Context) processFilters(p *PlacementPolicy) error {
	for _, f := range p.Filters() {
		if err := c.processFilter(f, true); err != nil {
			return err
		}
	}
	return nil
}

func (c *Context) processFilter(f *Filter, top bool) error {
	if f == nil {
		return fmt.Errorf("%w: FILTER", ErrMissingField)
	}
	if f.Name() == MainFilterName {
		return fmt.Errorf("%w: '*' is reserved", ErrInvalidFilterName)
	}
	if top && f.Name() == "" {
		return ErrUnnamedTopFilter
	}
	if !top && f.Name() != "" && c.Filters[f.Name()] == nil {
		return fmt.Errorf("%w: '%s'", ErrFilterNotFound, f.Name())
	}
	switch f.Operation() {
	case OpAND, OpOR:
		for _, flt := range f.InnerFilters() {
			if err := c.processFilter(flt, false); err != nil {
				return err
			}
		}
	default:
		if len(f.InnerFilters()) != 0 {
			return ErrNonEmptyFilters
		} else if !top && f.Name() != "" { // named reference
			return nil
		}
		switch f.Operation() {
		case OpEQ, OpNE:
		case OpGT, OpGE, OpLT, OpLE:
			n, err := strconv.ParseUint(f.Value(), 10, 64)
			if err != nil {
				return fmt.Errorf("%w: '%s'", ErrInvalidNumber, f.Value())
			}
			c.numCache[f] = n
		default:
			return fmt.Errorf("%w: %s", ErrInvalidFilterOp, f.Operation())
		}
	}
	if top {
		c.Filters[f.Name()] = f
	}
	return nil
}

// match matches f against b. It returns no errors because
// filter should have been parsed during context creation
// and missing node properties are considered as a regular fail.
func (c *Context) match(f *Filter, b *Node) bool {
	switch f.Operation() {
	case OpAND, OpOR:
		for _, lf := range f.InnerFilters() {
			if lf.Name() != "" {
				lf = c.Filters[lf.Name()]
			}
			ok := c.match(lf, b)
			if ok == (f.Operation() == OpOR) {
				return ok
			}
		}
		return f.Operation() == OpAND
	default:
		return c.matchKeyValue(f, b)
	}
}

func (c *Context) matchKeyValue(f *Filter, b *Node) bool {
	switch f.Operation() {
	case OpEQ:
		return b.Attribute(f.Key()) == f.Value()
	case OpNE:
		return b.Attribute(f.Key()) != f.Value()
	default:
		var attr uint64
		switch f.Key() {
		case PriceAttr:
			attr = b.Price
		case CapacityAttr:
			attr = b.Capacity
		default:
			var err error
			attr, err = strconv.ParseUint(b.Attribute(f.Key()), 10, 64)
			if err != nil {
				// Note: because filters are somewhat independent from nodes attributes,
				// We don't report an error here, and fail filter instead.
				return false
			}
		}
		switch f.Operation() {
		case OpGT:
			return attr > c.numCache[f]
		case OpGE:
			return attr >= c.numCache[f]
		case OpLT:
			return attr < c.numCache[f]
		case OpLE:
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

func filtersFromV2(fs []*netmap.Filter) []*Filter {
	res := make([]*Filter, 0, len(fs))

	for i := range fs {
		res = append(res, NewFilterFromV2(fs[i]))
	}

	return res
}

// InnerFilters returns list of inner filters.
func (f *Filter) InnerFilters() []*Filter {
	return filtersFromV2(
		(*netmap.Filter)(f).
			GetFilters(),
	)
}

func filtersToV2(fs []*Filter) []*netmap.Filter {
	fsV2 := make([]*netmap.Filter, 0, len(fs))

	for i := range fs {
		fsV2 = append(fsV2, fs[i].ToV2())
	}

	return fsV2
}

// SetInnerFilters sets list of inner filters.
func (f *Filter) SetInnerFilters(fs ...*Filter) {
	(*netmap.Filter)(f).
		SetFilters(filtersToV2(fs))
}

// Marshal marshals Filter into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (f *Filter) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*netmap.Filter)(f).StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Filter.
func (f *Filter) Unmarshal(data []byte) error {
	return (*netmap.Filter)(f).
		Unmarshal(data)
}

// MarshalJSON encodes Filter to protobuf JSON format.
func (f *Filter) MarshalJSON() ([]byte, error) {
	return (*netmap.Filter)(f).
		MarshalJSON()
}

// UnmarshalJSON decodes Filter from protobuf JSON format.
func (f *Filter) UnmarshalJSON(data []byte) error {
	return (*netmap.Filter)(f).
		UnmarshalJSON(data)
}
