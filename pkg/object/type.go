package object

import (
	"errors"

	"github.com/nspcc-dev/neofs-api-go/v2/object"
)

type Type uint8

const (
	TypeRegular Type = iota
	TypeTombstone
	TypeStorageGroup
)

func (t Type) ToV2() object.Type {
	if t2, ok := typeToV2(t); ok {
		return t2
	}

	return object.TypeRegular
}

// converts Action to v2 Action enum value. Returns false if value is not a named constant.
func typeToV2(t Type) (object.Type, bool) {
	switch t {
	default:
		return 0, false
	case TypeRegular:
		return object.TypeRegular, true
	case TypeTombstone:
		return object.TypeTombstone, true
	case TypeStorageGroup:
		return object.TypeStorageGroup, true
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

// String implements fmt.Stringer.
//
// Use MarshalText to get the canonical text format.
func (t Type) String() string {
	// TODO: simplify stringer after FromString will be removed (neofs-api-go#346)
	txt, _ := t.MarshalText()
	return string(txt)
}

var errUnsupportedType = errors.New("unsupported Type")

// MarshalText implements encoding.TextMarshaler.
//
// Text mapping:
//  * TypeTombstone: TOMBSTONE;
//  * TypeStorageGroup: STORAGE_GROUP;
//  * TypeRegular: REGULAR.
func (t Type) MarshalText() ([]byte, error) {
	t2, ok := typeToV2(t)
	if !ok {
		return nil, errUnsupportedType
	}

	return []byte(t2.String()), nil
}

func (t *Type) UnmarshalText(text []byte) error {
	var a2 object.Type

	ok := a2.FromString(string(text))
	if !ok {
		return errUnsupportedType
	}

	*t = TypeFromV2(a2)

	return nil
}

// FromString parses Type from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
//
// Deprecated: use UnmarshalText instead.
func (t *Type) FromString(s string) bool {
	return t.UnmarshalText([]byte(s)) == nil
}
