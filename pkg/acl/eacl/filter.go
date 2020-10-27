package eacl

import (
	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
)

// Filter defines check conditions if request header is matched or not. Matched
// header means that request should be processed according to EACL action.
type Filter struct {
	from    FilterHeaderType
	key     string
	matcher Match
	value   string
}

func (a Filter) Value() string {
	return a.value
}

func (a Filter) Matcher() Match {
	return a.matcher
}

func (a Filter) Key() string {
	return a.key
}

func (a Filter) From() FilterHeaderType {
	return a.from
}

func (a *Filter) ToV2() *v2acl.HeaderFilter {
	filter := new(v2acl.HeaderFilter)
	filter.SetValue(a.value)
	filter.SetKey(a.key)
	filter.SetMatchType(a.matcher.ToV2())
	filter.SetHeaderType(a.from.ToV2())

	return filter
}

func NewFilterFromV2(filter *v2acl.HeaderFilter) *Filter {
	f := new(Filter)

	if filter == nil {
		return f
	}

	f.from = FilterHeaderTypeFromV2(filter.GetHeaderType())
	f.matcher = MatchFromV2(filter.GetMatchType())
	f.key = filter.GetKey()
	f.value = filter.GetValue()

	return f
}
