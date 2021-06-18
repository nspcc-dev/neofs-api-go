package netmap

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

func TestClauseFromV2(t *testing.T) {
	for _, item := range []struct {
		c   Clause
		cV2 netmap.Clause
	}{
		{
			c:   ClauseUnspecified,
			cV2: netmap.UnspecifiedClause,
		},
		{
			c:   ClauseSame,
			cV2: netmap.Same,
		},
		{
			c:   ClauseDistinct,
			cV2: netmap.Distinct,
		},
	} {
		require.Equal(t, item.c, ClauseFromV2(item.cV2))
		require.Equal(t, item.cV2, item.c.ToV2())
	}
}

func TestClause_String(t *testing.T) {
	toPtr := func(v Clause) *Clause {
		return &v
	}

	testEnumStrings(t, new(Clause), []enumStringItem{
		{val: toPtr(ClauseDistinct), str: "DISTINCT"},
		{val: toPtr(ClauseSame), str: "SAME"},
		{val: toPtr(ClauseUnspecified), str: "CLAUSE_UNSPECIFIED"},
	})
}
