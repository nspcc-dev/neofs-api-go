package object

import (
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	v2object "github.com/nspcc-dev/neofs-api-go/v2/object"
)

// SearchMatchType indicates match operation on specified header.
type SearchMatchType uint32

const (
	MatchUnknown SearchMatchType = iota
	MatchStringEqual
)

func (m SearchMatchType) ToV2() v2object.MatchType {
	switch m {
	case MatchStringEqual:
		return v2object.MatchStringEqual
	default:
		return v2object.MatchUnknown
	}
}

func SearchMatchFromV2(t v2object.MatchType) (m SearchMatchType) {
	switch t {
	case v2object.MatchStringEqual:
		m = MatchStringEqual
	default:
		m = MatchUnknown
	}

	return m
}

type SearchFilter struct {
	header filterKey
	value  fmt.Stringer
	op     SearchMatchType
}

type staticStringer string

type filterKey struct {
	typ filterKeyType

	str string
}

// enumeration of reserved filter keys.
type filterKeyType int

type boolStringer bool

type SearchFilters []SearchFilter

const (
	_ filterKeyType = iota
	fKeyVersion
	fKeyObjectID
	fKeyContainerID
	fKeyOwnerID
	fKeyCreationEpoch
	fKeyPayloadLength
	fKeyPayloadHash
	fKeyType
	fKeyHomomorphicHash
	fKeyParent
	fKeyPropRoot
	fKeyPropPhy
	fKeyPropChildfree
)

func (k filterKey) String() string {
	switch k.typ {
	default:
		return k.str
	case fKeyVersion:
		return v2object.FilterHeaderVersion
	case fKeyObjectID:
		return v2object.FilterHeaderObjectID
	case fKeyContainerID:
		return v2object.FilterHeaderContainerID
	case fKeyOwnerID:
		return v2object.FilterHeaderOwnerID
	case fKeyCreationEpoch:
		return v2object.FilterHeaderCreationEpoch
	case fKeyPayloadLength:
		return v2object.FilterHeaderPayloadLength
	case fKeyPayloadHash:
		return v2object.FilterHeaderPayloadHash
	case fKeyType:
		return v2object.FilterHeaderObjectType
	case fKeyHomomorphicHash:
		return v2object.FilterHeaderHomomorphicHash
	case fKeyParent:
		return v2object.FilterHeaderParent
	case fKeyPropRoot:
		return v2object.FilterPropertyRoot
	case fKeyPropPhy:
		return v2object.FilterPropertyPhy
	case fKeyPropChildfree:
		return v2object.FilterPropertyChildfree
	}
}

func (s staticStringer) String() string {
	return string(s)
}

func (s boolStringer) String() string {
	if s {
		return v2object.BooleanPropertyValueTrue
	}

	return v2object.BooleanPropertyValueFalse
}

func (f *SearchFilter) Header() string {
	return f.header.String()
}

func (f *SearchFilter) Value() string {
	return f.value.String()
}

func (f *SearchFilter) Operation() SearchMatchType {
	return f.op
}

func NewSearchFilters() SearchFilters {
	return SearchFilters{}
}

func NewSearchFiltersFromV2(v2 []*v2object.SearchFilter) SearchFilters {
	filters := make(SearchFilters, 0, len(v2))
	for i := range v2 {
		if v2[i] == nil {
			continue
		}

		filters.AddFilter(
			v2[i].GetKey(),
			v2[i].GetValue(),
			SearchMatchFromV2(v2[i].GetMatchType()),
		)
	}

	return filters
}

func (f *SearchFilters) addFilter(op SearchMatchType, keyTyp filterKeyType, key string, val fmt.Stringer) {
	if *f == nil {
		*f = make(SearchFilters, 0, 1)
	}

	*f = append(*f, SearchFilter{
		header: filterKey{
			typ: keyTyp,
			str: key,
		},
		value: val,
		op:    op,
	})
}

func (f *SearchFilters) AddFilter(header, value string, op SearchMatchType) {
	f.addFilter(op, 0, header, staticStringer(value))
}

func (f *SearchFilters) addReservedFilter(op SearchMatchType, keyTyp filterKeyType, val fmt.Stringer) {
	f.addFilter(op, keyTyp, "", val)
}

// addFlagFilters adds filters that works like flags: they don't need to have
// specific match type or value. They processed by NeoFS nodes by the fact
// of presence in search query. E.g.: PHY, ROOT.
func (f *SearchFilters) addFlagFilter(keyTyp filterKeyType) {
	f.addFilter(MatchUnknown, keyTyp, "", staticStringer(""))
}

func (f *SearchFilters) AddObjectVersionFilter(op SearchMatchType, v *pkg.Version) {
	f.addReservedFilter(op, fKeyVersion, v)
}

func (f *SearchFilters) AddObjectContainerIDFilter(m SearchMatchType, id *container.ID) {
	f.addReservedFilter(m, fKeyContainerID, id)
}

func (f *SearchFilters) AddObjectOwnerIDFilter(m SearchMatchType, id *owner.ID) {
	f.addReservedFilter(m, fKeyOwnerID, id)
}

func (f SearchFilters) ToV2() []*v2object.SearchFilter {
	result := make([]*v2object.SearchFilter, 0, len(f))
	for i := range f {
		v2 := new(v2object.SearchFilter)
		v2.SetKey(f[i].header.String())
		v2.SetValue(f[i].value.String())
		v2.SetMatchType(f[i].op.ToV2())

		result = append(result, v2)
	}

	return result
}

func (f *SearchFilters) addRootFilter() {
	f.addFlagFilter(fKeyPropRoot)
}

func (f *SearchFilters) AddRootFilter() {
	f.addRootFilter()
}

func (f *SearchFilters) addPhyFilter() {
	f.addFlagFilter(fKeyPropPhy)
}

func (f *SearchFilters) AddPhyFilter() {
	f.addPhyFilter()
}

func (f *SearchFilters) addChildFreeFilter(val bool) {
	f.addReservedFilter(MatchStringEqual, fKeyPropChildfree, boolStringer(val))
}

func (f *SearchFilters) AddChildfreeFilter() {
	f.addChildFreeFilter(true)
}

func (f *SearchFilters) AddNonChildfreeFilter() {
	f.addChildFreeFilter(false)
}

func (f *SearchFilters) AddParentIDFilter(m SearchMatchType, id *ID) {
	f.addReservedFilter(m, fKeyParent, id)
}
