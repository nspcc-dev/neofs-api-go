package netmap

import (
	"errors"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// Clause is an enumeration of selector modifiers
// that shows how the node set will be formed.
type Clause uint32

const (
	ClauseUnspecified Clause = iota

	// ClauseSame is a selector modifier to select only nodes having the same value of bucket attribute.
	ClauseSame

	// ClauseDistinct is a selector modifier to select nodes having different values of bucket attribute.
	ClauseDistinct
)

// ClauseFromV2 converts v2 Clause to Clause.
func ClauseFromV2(c netmap.Clause) Clause {
	switch c {
	default:
		return ClauseUnspecified
	case netmap.Same:
		return ClauseSame
	case netmap.Distinct:
		return ClauseDistinct
	}
}

// ToV2 converts Clause to v2 Clause.
func (c Clause) ToV2() netmap.Clause {
	if o2, ok := clauseToV2(c); ok {
		return o2
	}

	return netmap.UnspecifiedClause
}

// converts Clause to v2 Clause. enum value. Returns false if value is not a named constant.
func clauseToV2(c Clause) (netmap.Clause, bool) {
	switch c {
	default:
		return 0, false
	case ClauseUnspecified:
		return netmap.UnspecifiedClause, true
	case ClauseDistinct:
		return netmap.Distinct, true
	case ClauseSame:
		return netmap.Same, true
	}
}

// String implements fmt.Stringer.
//
// Use MarshalText to get the canonical text format.
func (c Clause) String() string {
	// TODO: simplify stringer after FromString will be removed (neofs-api-go#346)
	txt, _ := c.MarshalText()
	return string(txt)
}

var errUnsupportedClause = errors.New("unsupported Clause")

// MarshalText implements encoding.TextMarshaler.
//
// Text mapping:
//  * ClauseDistinct: DISTINCT;
//  * ClauseSame: SAME;
//  * ClauseUnspecified: CLAUSE_UNSPECIFIED.
func (c Clause) MarshalText() ([]byte, error) {
	o2, ok := clauseToV2(c)
	if !ok {
		return nil, errUnsupportedClause
	}

	return []byte(o2.String()), nil
}

func (c *Clause) UnmarshalText(text []byte) error {
	var c2 netmap.Clause

	ok := c2.FromString(string(text))
	if !ok {
		return errUnsupportedClause
	}

	*c = ClauseFromV2(c2)

	return nil
}

// FromString parses Clause from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
//
// Deprecated: use UnmarshalText instead.
func (c *Clause) FromString(s string) bool {
	return c.UnmarshalText([]byte(s)) == nil
}
