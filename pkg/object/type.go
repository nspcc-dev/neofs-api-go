package object

import (
	"github.com/nspcc-dev/neofs-api-go/v2/object"
)

type Type uint8

const (
	TypeRegular Type = iota
	TypeTombstone
	TypeStorageGroup
)

func (t Type) ToV2() object.Type {
	switch t {
	case TypeTombstone:
		return object.TypeTombstone
	case TypeStorageGroup:
		return object.TypeStorageGroup
	default:
		return object.TypeRegular
	}
}

func TypeFromV2(t object.Type) Type {
	switch t {
	case object.TypeTombstone:
		return TypeTombstone
	case object.TypeStorageGroup:
		return TypeStorageGroup
	default:
		return TypeRegular
	}
}

func (t Type) String() string {
	return t.ToV2().String()
}

// TypeFromString parses Type from its string representation.
func TypeFromString(s string) Type {
	return TypeFromV2(
		object.TypeFromString(s),
	)
}
