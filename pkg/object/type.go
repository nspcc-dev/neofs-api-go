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

// String returns string representation of Type.
//
// String mapping:
//  * TypeTombstone: TOMBSTONE;
//  * TypeStorageGroup: STORAGE_GROUP;
//  * TypeRegular, default: REGULAR.
func (t Type) String() string {
	return t.ToV2().String()
}

// FromString parses Type from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (t *Type) FromString(s string) bool {
	var g object.Type

	ok := g.FromString(s)

	if ok {
		*t = TypeFromV2(g)
	}

	return ok
}

// TypeFromString parses Type from its string representation.
//
// Deprecated: use FromString method.
func TypeFromString(s string) (t Type) {
	t.FromString(s)
	return
}
