package eacl

import (
	"fmt"

	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
)

// Filter defines check conditions if request header is matched or not. Matched
// header means that request should be processed according to EACL action.
type Filter struct {
	from    FilterHeaderType
	key     filterKey
	matcher Match
	value   fmt.Stringer
}

type staticStringer string

type filterKey struct {
	typ filterKeyType

	str string
}

// enumeration of reserved filter keys.
type filterKeyType int

const (
	_ filterKeyType = iota
	fKeyObjVersion
	fKeyObjID
	fKeyObjContainerID
	fKeyObjOwnerID
	fKeyObjCreationEpoch
	fKeyObjPayloadLength
	fKeyObjPayloadHash
	fKeyObjType
	fKeyObjHomomorphicHash
)

func (s staticStringer) String() string {
	return string(s)
}

func (a Filter) Value() string {
	return a.value.String()
}

func (a Filter) Matcher() Match {
	return a.matcher
}

func (a Filter) Key() string {
	return a.key.String()
}

func (a Filter) From() FilterHeaderType {
	return a.from
}

func (a *Filter) ToV2() *v2acl.HeaderFilter {
	filter := new(v2acl.HeaderFilter)
	filter.SetValue(a.value.String())
	filter.SetKey(a.key.String())
	filter.SetMatchType(a.matcher.ToV2())
	filter.SetHeaderType(a.from.ToV2())

	return filter
}

func (k filterKey) String() string {
	switch k.typ {
	default:
		return k.str
	case fKeyObjVersion:
		return v2acl.FilterObjectVersion
	case fKeyObjID:
		return v2acl.FilterObjectID
	case fKeyObjContainerID:
		return v2acl.FilterObjectContainerID
	case fKeyObjOwnerID:
		return v2acl.FilterObjectOwnerID
	case fKeyObjCreationEpoch:
		return v2acl.FilterObjectCreationEpoch
	case fKeyObjPayloadLength:
		return v2acl.FilterObjectPayloadLength
	case fKeyObjPayloadHash:
		return v2acl.FilterObjectPayloadHash
	case fKeyObjType:
		return v2acl.FilterObjectType
	case fKeyObjHomomorphicHash:
		return v2acl.FilterObjectHomomorphicHash
	}
}

func NewFilterFromV2(filter *v2acl.HeaderFilter) *Filter {
	f := new(Filter)

	if filter == nil {
		return f
	}

	f.from = FilterHeaderTypeFromV2(filter.GetHeaderType())
	f.matcher = MatchFromV2(filter.GetMatchType())
	f.key.str = filter.GetKey()
	f.value = staticStringer(filter.GetValue())

	return f
}
