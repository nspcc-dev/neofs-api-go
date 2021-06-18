package netmap

import (
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
	switch c {
	default:
		return netmap.UnspecifiedClause
	case ClauseDistinct:
		return netmap.Distinct
	case ClauseSame:
		return netmap.Same
	}
}

// String returns string representation of Clause.
//
// String mapping:
//  * ClauseDistinct: DISTINCT;
//  * ClauseSame: SAME;
//  * ClauseUnspecified, default: CLAUSE_UNSPECIFIED.
func (c Clause) String() string {
	return c.ToV2().String()
}

// FromString parses Clause from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (c *Clause) FromString(s string) bool {
	var g netmap.Clause

	ok := g.FromString(s)

	if ok {
		*c = ClauseFromV2(g)
	}

	return ok
}
