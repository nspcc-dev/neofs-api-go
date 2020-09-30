package object_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	v2object "github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/stretchr/testify/require"
)

var (
	eqV2Matches = map[object.SearchMatchType]v2object.MatchType{
		object.MatchUnknown:     v2object.MatchUnknown,
		object.MatchStringEqual: v2object.MatchStringEqual,
	}
)

func TestMatch(t *testing.T) {
	t.Run("known matches", func(t *testing.T) {
		for i := object.MatchUnknown; i <= object.MatchStringEqual; i++ {
			require.Equal(t, eqV2Matches[i], i.ToV2())
			require.Equal(t, object.SearchMatchFromV2(i.ToV2()), i)
		}
	})

	t.Run("unknown matches", func(t *testing.T) {
		require.Equal(t, (object.MatchStringEqual + 1).ToV2(), v2object.MatchUnknown)
		require.Equal(t, object.SearchMatchFromV2(v2object.MatchStringEqual+1), object.MatchUnknown)
	})
}

func TestFilter(t *testing.T) {
	inputs := [][]string{
		{"user-header", "user-value"},
		{object.HdrSysNameID, "objectID"},
	}

	filters := object.NewSearchFilters()
	for i := range inputs {
		filters.AddFilter(inputs[i][0], inputs[i][1], object.MatchStringEqual)
	}

	require.Len(t, filters, len(inputs))
	for i := range inputs {
		require.Equal(t, inputs[i][0], filters[i].Header())
		require.Equal(t, inputs[i][1], filters[i].Value())
		require.Equal(t, object.MatchStringEqual, filters[i].Operation())
	}

	v2 := filters.ToV2()
	newFilters := object.NewSearchFiltersFromV2(v2)
	require.Equal(t, filters, newFilters)
}

func TestSearchFilters_AddRootFilter(t *testing.T) {
	fs := new(object.SearchFilters)

	fs.AddRootFilter()

	require.Len(t, *fs, 1)

	f := (*fs)[0]

	require.Equal(t, object.MatchStringEqual, f.Operation())
	require.Equal(t, object.KeyRoot, f.Header())
	require.Equal(t, object.ValRoot, f.Value())
}

func TestSearchFilters_AddNonRootFilter(t *testing.T) {
	fs := new(object.SearchFilters)

	fs.AddNonRootFilter()

	require.Len(t, *fs, 1)

	f := (*fs)[0]

	require.Equal(t, object.MatchStringEqual, f.Operation())
	require.Equal(t, object.KeyRoot, f.Header())
	require.Equal(t, object.ValNonRoot, f.Value())
}

func TestSearchFilters_AddLeafFilter(t *testing.T) {
	fs := new(object.SearchFilters)

	fs.AddLeafFilter()

	require.Len(t, *fs, 1)

	f := (*fs)[0]

	require.Equal(t, object.MatchStringEqual, f.Operation())
	require.Equal(t, object.KeyLeaf, f.Header())
	require.Equal(t, object.ValLeaf, f.Value())
}

func TestSearchFilters_AddNonLeafFilter(t *testing.T) {
	fs := new(object.SearchFilters)

	fs.AddNonLeafFilter()

	require.Len(t, *fs, 1)

	f := (*fs)[0]

	require.Equal(t, object.MatchStringEqual, f.Operation())
	require.Equal(t, object.KeyLeaf, f.Header())
	require.Equal(t, object.ValNonLeaf, f.Value())
}
