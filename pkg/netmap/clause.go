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

func (c Clause) String() string {
	switch c {
	default:
		return "UNSPECIFIED"
	case ClauseDistinct:
		return "DISTINCT"
	case ClauseSame:
		return "SAME"
	}
}
