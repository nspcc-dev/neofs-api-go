package object

import (
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
	header string
	value  string
	op     SearchMatchType
}

type SearchFilters []SearchFilter

func (f *SearchFilter) Header() string {
	return f.header
}

func (f *SearchFilter) Value() string {
	return f.value
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

		filters = append(filters, SearchFilter{
			header: v2[i].GetName(),
			value:  v2[i].GetValue(),
			op:     SearchMatchFromV2(v2[i].GetMatchType()),
		})
	}

	return filters
}

func (f *SearchFilters) AddFilter(header, value string, op SearchMatchType) {
	if *f == nil {
		*f = make(SearchFilters, 0, 1)
	}

	*f = append(*f, SearchFilter{
		header: header,
		value:  value,
		op:     op,
	})
}

func (f SearchFilters) ToV2() []*v2object.SearchFilter {
	result := make([]*v2object.SearchFilter, 0, len(f))
	for i := range f {
		v2 := new(v2object.SearchFilter)
		v2.SetName(f[i].header)
		v2.SetValue(f[i].value)
		v2.SetMatchType(f[i].op.ToV2())

		result = append(result, v2)
	}

	return result
}

func (f *SearchFilters) addRootFilter(val string) {
	f.AddFilter(KeyRoot, val, MatchStringEqual)
}

func (f *SearchFilters) AddRootFilter() {
	f.addRootFilter(ValRoot)
}

func (f *SearchFilters) AddNonRootFilter() {
	f.addRootFilter(ValNonRoot)
}
