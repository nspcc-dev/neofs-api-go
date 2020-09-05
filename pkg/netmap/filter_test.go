package netmap

import (
	"errors"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

func TestContext_ProcessFilters(t *testing.T) {
	fs := []*netmap.Filter{
		newFilter("StorageSSD", "Storage", "SSD", netmap.EQ),
		newFilter("GoodRating", "Rating", "4", netmap.GE),
		newFilter("Main", "", "", netmap.AND,
			newFilter("StorageSSD", "", "", 0),
			newFilter("", "IntField", "123", netmap.LT),
			newFilter("GoodRating", "", "", 0)),
	}
	nm, err := NewNetmap(nil)
	require.NoError(t, err)
	c := NewContext(nm)
	p := newPlacementPolicy(1, nil, nil, fs)
	require.NoError(t, c.processFilters(p))
	require.Equal(t, 3, len(c.Filters))
	for _, f := range fs {
		require.Equal(t, f, c.Filters[f.GetName()])
	}

	require.Equal(t, uint64(4), c.numCache[fs[1]])
	require.Equal(t, uint64(123), c.numCache[fs[2].GetFilters()[1]])
}

func TestContext_ProcessFiltersInvalid(t *testing.T) {
	errTestCases := []struct {
		name   string
		filter *netmap.Filter
		err    error
	}{
		{
			"UnnamedTop",
			newFilter("", "Storage", "SSD", netmap.EQ),
			ErrUnnamedTopFilter,
		},
		{
			"InvalidReference",
			newFilter("Main", "", "", netmap.AND,
				newFilter("StorageSSD", "", "", 0)),
			ErrFilterNotFound,
		},
		{
			"NonEmptyKeyed",
			newFilter("Main", "Storage", "SSD", netmap.EQ,
				newFilter("StorageSSD", "", "", 0)),
			ErrNonEmptyFilters,
		},
		{
			"InvalidNumber",
			newFilter("Main", "Rating", "three", netmap.GE),
			ErrInvalidNumber,
		},
		{
			"InvalidOp",
			newFilter("Main", "Rating", "3", netmap.UnspecifiedOperation),
			ErrInvalidFilterOp,
		},
		{
			"InvalidName",
			newFilter("*", "Rating", "3", netmap.GE),
			ErrInvalidFilterName,
		},
		{
			"MissingFilter",
			nil,
			ErrMissingField,
		},
	}
	for _, tc := range errTestCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewContext(new(Netmap))
			p := newPlacementPolicy(1, nil, nil, []*netmap.Filter{tc.filter})
			err := c.processFilters(p)
			require.True(t, errors.Is(err, tc.err), "got: %v", err)
		})
	}
}

func TestFilter_MatchSimple(t *testing.T) {
	b := &Node{AttrMap: map[string]string{
		"Rating":  "4",
		"Country": "Germany",
	}}
	testCases := []struct {
		name string
		ok   bool
		f    *netmap.Filter
	}{
		{
			"GE_true", true,
			newFilter("Main", "Rating", "4", netmap.GE),
		},
		{
			"GE_false", false,
			newFilter("Main", "Rating", "5", netmap.GE),
		},
		{
			"GT_true", true,
			newFilter("Main", "Rating", "3", netmap.GT),
		},
		{
			"GT_false", false,
			newFilter("Main", "Rating", "4", netmap.GT),
		},
		{
			"LE_true", true,
			newFilter("Main", "Rating", "4", netmap.LE),
		},
		{
			"LE_false", false,
			newFilter("Main", "Rating", "3", netmap.LE),
		},
		{
			"LT_true", true,
			newFilter("Main", "Rating", "5", netmap.LT),
		},
		{
			"LT_false", false,
			newFilter("Main", "Rating", "4", netmap.LT),
		},
		{
			"EQ_true", true,
			newFilter("Main", "Country", "Germany", netmap.EQ),
		},
		{
			"EQ_false", false,
			newFilter("Main", "Country", "China", netmap.EQ),
		},
		{
			"NE_true", true,
			newFilter("Main", "Country", "France", netmap.NE),
		},
		{
			"NE_false", false,
			newFilter("Main", "Country", "Germany", netmap.NE),
		},
	}
	for _, tc := range testCases {
		c := NewContext(new(Netmap))
		p := newPlacementPolicy(1, nil, nil, []*netmap.Filter{tc.f})
		require.NoError(t, c.processFilters(p))
		require.Equal(t, tc.ok, c.match(tc.f, b))
	}

	t.Run("InvalidOp", func(t *testing.T) {
		f := newFilter("Main", "Rating", "5", netmap.EQ)
		c := NewContext(new(Netmap))
		p := newPlacementPolicy(1, nil, nil, []*netmap.Filter{f})
		require.NoError(t, c.processFilters(p))

		// just for the coverage
		f.SetOp(netmap.UnspecifiedOperation)
		require.False(t, c.match(f, b))
	})
}

func TestFilter_Match(t *testing.T) {
	fs := []*netmap.Filter{
		newFilter("StorageSSD", "Storage", "SSD", netmap.EQ),
		newFilter("GoodRating", "Rating", "4", netmap.GE),
		newFilter("Main", "", "", netmap.AND,
			newFilter("StorageSSD", "", "", 0),
			newFilter("", "IntField", "123", netmap.LT),
			newFilter("GoodRating", "", "", 0),
			newFilter("", "", "", netmap.OR,
				newFilter("", "Param", "Value1", netmap.EQ),
				newFilter("", "Param", "Value2", netmap.EQ),
			)),
	}
	c := NewContext(new(Netmap))
	p := newPlacementPolicy(1, nil, nil, fs)
	require.NoError(t, c.processFilters(p))

	t.Run("Good", func(t *testing.T) {
		n := getTestNode("Storage", "SSD", "Rating", "10", "IntField", "100", "Param", "Value1")
		require.True(t, c.applyFilter("Main", n))
	})
	t.Run("InvalidStorage", func(t *testing.T) {
		n := getTestNode("Storage", "HDD", "Rating", "10", "IntField", "100", "Param", "Value1")
		require.False(t, c.applyFilter("Main", n))
	})
	t.Run("InvalidRating", func(t *testing.T) {
		n := getTestNode("Storage", "SSD", "Rating", "3", "IntField", "100", "Param", "Value1")
		require.False(t, c.applyFilter("Main", n))
	})
	t.Run("InvalidIntField", func(t *testing.T) {
		n := getTestNode("Storage", "SSD", "Rating", "3", "IntField", "str", "Param", "Value1")
		require.False(t, c.applyFilter("Main", n))
	})
	t.Run("InvalidParam", func(t *testing.T) {
		n := getTestNode("Storage", "SSD", "Rating", "3", "IntField", "100", "Param", "NotValue")
		require.False(t, c.applyFilter("Main", n))
	})
}
