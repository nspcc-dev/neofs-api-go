package netmap

import (
	"errors"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

func TestContext_ProcessFilters(t *testing.T) {
	fs := []*Filter{
		newFilter("StorageSSD", "Storage", "SSD", OpEQ),
		newFilter("GoodRating", "Rating", "4", OpGE),
		newFilter("Main", "", "", OpAND,
			newFilter("StorageSSD", "", "", 0),
			newFilter("", "IntField", "123", OpLT),
			newFilter("GoodRating", "", "", 0)),
	}
	nm, err := NewNetmap(nil)
	require.NoError(t, err)
	c := NewContext(nm)
	p := newPlacementPolicy(1, nil, nil, fs)
	require.NoError(t, c.processFilters(p))
	require.Equal(t, 3, len(c.Filters))
	for _, f := range fs {
		require.Equal(t, f, c.Filters[f.Name()])
	}

	require.Equal(t, uint64(4), c.numCache[fs[1]])
	require.Equal(t, uint64(123), c.numCache[fs[2].InnerFilters()[1]])
}

func TestContext_ProcessFiltersInvalid(t *testing.T) {
	errTestCases := []struct {
		name   string
		filter *Filter
		err    error
	}{
		{
			"UnnamedTop",
			newFilter("", "Storage", "SSD", OpEQ),
			ErrUnnamedTopFilter,
		},
		{
			"InvalidReference",
			newFilter("Main", "", "", OpAND,
				newFilter("StorageSSD", "", "", 0)),
			ErrFilterNotFound,
		},
		{
			"NonEmptyKeyed",
			newFilter("Main", "Storage", "SSD", OpEQ,
				newFilter("StorageSSD", "", "", 0)),
			ErrNonEmptyFilters,
		},
		{
			"InvalidNumber",
			newFilter("Main", "Rating", "three", OpGE),
			ErrInvalidNumber,
		},
		{
			"InvalidOp",
			newFilter("Main", "Rating", "3", 0),
			ErrInvalidFilterOp,
		},
		{
			"InvalidName",
			newFilter("*", "Rating", "3", OpGE),
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
			p := newPlacementPolicy(1, nil, nil, []*Filter{tc.filter})
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
		f    *Filter
	}{
		{
			"GE_true", true,
			newFilter("Main", "Rating", "4", OpGE),
		},
		{
			"GE_false", false,
			newFilter("Main", "Rating", "5", OpGE),
		},
		{
			"GT_true", true,
			newFilter("Main", "Rating", "3", OpGT),
		},
		{
			"GT_false", false,
			newFilter("Main", "Rating", "4", OpGT),
		},
		{
			"LE_true", true,
			newFilter("Main", "Rating", "4", OpLE),
		},
		{
			"LE_false", false,
			newFilter("Main", "Rating", "3", OpLE),
		},
		{
			"LT_true", true,
			newFilter("Main", "Rating", "5", OpLT),
		},
		{
			"LT_false", false,
			newFilter("Main", "Rating", "4", OpLT),
		},
		{
			"EQ_true", true,
			newFilter("Main", "Country", "Germany", OpEQ),
		},
		{
			"EQ_false", false,
			newFilter("Main", "Country", "China", OpEQ),
		},
		{
			"NE_true", true,
			newFilter("Main", "Country", "France", OpNE),
		},
		{
			"NE_false", false,
			newFilter("Main", "Country", "Germany", OpNE),
		},
	}
	for _, tc := range testCases {
		c := NewContext(new(Netmap))
		p := newPlacementPolicy(1, nil, nil, []*Filter{tc.f})
		require.NoError(t, c.processFilters(p))
		require.Equal(t, tc.ok, c.match(tc.f, b))
	}

	t.Run("InvalidOp", func(t *testing.T) {
		f := newFilter("Main", "Rating", "5", OpEQ)
		c := NewContext(new(Netmap))
		p := newPlacementPolicy(1, nil, nil, []*Filter{f})
		require.NoError(t, c.processFilters(p))

		// just for the coverage
		f.SetOperation(0)
		require.False(t, c.match(f, b))
	})
}

func TestFilter_Match(t *testing.T) {
	fs := []*Filter{
		newFilter("StorageSSD", "Storage", "SSD", OpEQ),
		newFilter("GoodRating", "Rating", "4", OpGE),
		newFilter("Main", "", "", OpAND,
			newFilter("StorageSSD", "", "", 0),
			newFilter("", "IntField", "123", OpLT),
			newFilter("GoodRating", "", "", 0),
			newFilter("", "", "", OpOR,
				newFilter("", "Param", "Value1", OpEQ),
				newFilter("", "Param", "Value2", OpEQ),
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

func testFilter() *Filter {
	f := NewFilter()
	f.SetOperation(OpGE)
	f.SetName("name")
	f.SetKey("key")
	f.SetValue("value")

	return f
}

func TestFilterFromV2(t *testing.T) {
	t.Run("nil from V2", func(t *testing.T) {
		var x *netmap.Filter

		require.Nil(t, NewFilterFromV2(x))
	})

	t.Run("nil to V2", func(t *testing.T) {
		var x *Filter

		require.Nil(t, x.ToV2())
	})

	fV2 := new(netmap.Filter)
	fV2.SetOp(netmap.GE)
	fV2.SetName("name")
	fV2.SetKey("key")
	fV2.SetValue("value")

	f := NewFilterFromV2(fV2)

	require.Equal(t, fV2, f.ToV2())
}

func TestFilter_Key(t *testing.T) {
	f := NewFilter()
	key := "some key"

	f.SetKey(key)

	require.Equal(t, key, f.Key())
}

func TestFilter_Value(t *testing.T) {
	f := NewFilter()
	val := "some value"

	f.SetValue(val)

	require.Equal(t, val, f.Value())
}

func TestFilter_Name(t *testing.T) {
	f := NewFilter()
	name := "some name"

	f.SetName(name)

	require.Equal(t, name, f.Name())
}

func TestFilter_Operation(t *testing.T) {
	f := NewFilter()
	op := OpGE

	f.SetOperation(op)

	require.Equal(t, op, f.Operation())
}

func TestFilter_InnerFilters(t *testing.T) {
	f := NewFilter()

	f1, f2 := testFilter(), testFilter()

	f.SetInnerFilters(f1, f2)

	require.Equal(t, []*Filter{f1, f2}, f.InnerFilters())
}

func TestFilterEncoding(t *testing.T) {
	f := newFilter("name", "key", "value", OpEQ,
		newFilter("name2", "key2", "value", OpOR),
	)

	t.Run("binary", func(t *testing.T) {
		data, err := f.Marshal()
		require.NoError(t, err)

		f2 := NewFilter()
		require.NoError(t, f2.Unmarshal(data))

		require.Equal(t, f, f2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := f.MarshalJSON()
		require.NoError(t, err)

		f2 := NewFilter()
		require.NoError(t, f2.UnmarshalJSON(data))

		require.Equal(t, f, f2)
	})
}

func TestNewFilter(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		filter := NewFilter()

		// check initial values
		require.Empty(t, filter.Name())
		require.Empty(t, filter.Key())
		require.Empty(t, filter.Value())
		require.Zero(t, filter.Operation())
		require.Nil(t, filter.InnerFilters())

		// convert to v2 message
		filterV2 := filter.ToV2()

		require.Empty(t, filterV2.GetName())
		require.Empty(t, filterV2.GetKey())
		require.Empty(t, filterV2.GetValue())
		require.Equal(t, netmap.UnspecifiedOperation, filterV2.GetOp())
		require.Nil(t, filterV2.GetFilters())
	})
}
