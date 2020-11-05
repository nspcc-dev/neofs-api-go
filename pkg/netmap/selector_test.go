package netmap

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

func TestPlacementPolicy_UnspecifiedClause(t *testing.T) {
	p := newPlacementPolicy(1,
		[]*Replica{newReplica(1, "X")},
		[]*Selector{
			newSelector("X", "", ClauseDistinct, 4, "*"),
		},
		nil,
	)
	nodes := []NodeInfo{
		nodeInfoFromAttributes("ID", "1", "Country", "RU", "City", "St.Petersburg", "SSD", "0"),
		nodeInfoFromAttributes("ID", "2", "Country", "RU", "City", "St.Petersburg", "SSD", "1"),
		nodeInfoFromAttributes("ID", "3", "Country", "RU", "City", "Moscow", "SSD", "1"),
		nodeInfoFromAttributes("ID", "4", "Country", "RU", "City", "Moscow", "SSD", "1"),
	}

	nm, err := NewNetmap(NodesFromInfo(nodes))
	require.NoError(t, err)
	v, err := nm.GetContainerNodes(p, nil)
	require.NoError(t, err)
	require.Equal(t, 4, len(v.Flatten()))
}

func TestPlacementPolicy_GetPlacementVectors(t *testing.T) {
	p := newPlacementPolicy(2,
		[]*Replica{
			newReplica(1, "SPB"),
			newReplica(2, "Americas"),
		},
		[]*Selector{
			newSelector("SPB", "City", ClauseSame, 1, "SPBSSD"),
			newSelector("Americas", "City", ClauseDistinct, 2, "Americas"),
		},
		[]*Filter{
			newFilter("SPBSSD", "", "", OpAND,
				newFilter("", "Country", "RU", OpEQ),
				newFilter("", "City", "St.Petersburg", OpEQ),
				newFilter("", "SSD", "1", OpEQ)),
			newFilter("Americas", "", "", OpOR,
				newFilter("", "Continent", "NA", OpEQ),
				newFilter("", "Continent", "SA", OpEQ)),
		})
	nodes := []NodeInfo{
		nodeInfoFromAttributes("ID", "1", "Country", "RU", "City", "St.Petersburg", "SSD", "0"),
		nodeInfoFromAttributes("ID", "2", "Country", "RU", "City", "St.Petersburg", "SSD", "1"),
		nodeInfoFromAttributes("ID", "3", "Country", "RU", "City", "Moscow", "SSD", "1"),
		nodeInfoFromAttributes("ID", "4", "Country", "RU", "City", "Moscow", "SSD", "1"),
		nodeInfoFromAttributes("ID", "5", "Country", "RU", "City", "St.Petersburg", "SSD", "1"),
		nodeInfoFromAttributes("ID", "6", "Continent", "NA", "City", "NewYork"),
		nodeInfoFromAttributes("ID", "7", "Continent", "AF", "City", "Cairo"),
		nodeInfoFromAttributes("ID", "8", "Continent", "AF", "City", "Cairo"),
		nodeInfoFromAttributes("ID", "9", "Continent", "SA", "City", "Lima"),
		nodeInfoFromAttributes("ID", "10", "Continent", "AF", "City", "Cairo"),
		nodeInfoFromAttributes("ID", "11", "Continent", "NA", "City", "NewYork"),
		nodeInfoFromAttributes("ID", "12", "Continent", "NA", "City", "LosAngeles"),
		nodeInfoFromAttributes("ID", "13", "Continent", "SA", "City", "Lima"),
	}

	nm, err := NewNetmap(NodesFromInfo(nodes))
	require.NoError(t, err)
	v, err := nm.GetContainerNodes(p, nil)
	require.NoError(t, err)
	require.Equal(t, 2, len(v.Replicas()))
	require.Equal(t, 6, len(v.Flatten()))

	require.Equal(t, 2, len(v.Replicas()[0]))
	ids := map[string]struct{}{}
	for _, ni := range v.Replicas()[0] {
		require.Equal(t, "RU", ni.Attribute("Country"))
		require.Equal(t, "St.Petersburg", ni.Attribute("City"))
		require.Equal(t, "1", ni.Attribute("SSD"))
		ids[ni.Attribute("ID")] = struct{}{}
	}
	require.Equal(t, len(v.Replicas()[0]), len(ids), "not all nodes we distinct")

	require.Equal(t, 4, len(v.Replicas()[1])) // 2 cities * 2 HRWB
	ids = map[string]struct{}{}
	for _, ni := range v.Replicas()[1] {
		require.Contains(t, []string{"NA", "SA"}, ni.Attribute("Continent"))
		ids[ni.Attribute("ID")] = struct{}{}
	}
	require.Equal(t, len(v.Replicas()[1]), len(ids), "not all nodes we distinct")
}

func TestPlacementPolicy_ProcessSelectors(t *testing.T) {
	p := newPlacementPolicy(2, nil,
		[]*Selector{
			newSelector("SameRU", "City", ClauseSame, 2, "FromRU"),
			newSelector("DistinctRU", "City", ClauseDistinct, 2, "FromRU"),
			newSelector("Good", "Country", ClauseDistinct, 2, "Good"),
			newSelector("Main", "Country", ClauseDistinct, 3, "*"),
		},
		[]*Filter{
			newFilter("FromRU", "Country", "Russia", OpEQ),
			newFilter("Good", "Rating", "4", OpGE),
		})
	nodes := []NodeInfo{
		nodeInfoFromAttributes("Country", "Russia", "Rating", "1", "City", "SPB"),
		nodeInfoFromAttributes("Country", "Germany", "Rating", "5", "City", "Berlin"),
		nodeInfoFromAttributes("Country", "Russia", "Rating", "6", "City", "Moscow"),
		nodeInfoFromAttributes("Country", "France", "Rating", "4", "City", "Paris"),
		nodeInfoFromAttributes("Country", "France", "Rating", "1", "City", "Lyon"),
		nodeInfoFromAttributes("Country", "Russia", "Rating", "5", "City", "SPB"),
		nodeInfoFromAttributes("Country", "Russia", "Rating", "7", "City", "Moscow"),
		nodeInfoFromAttributes("Country", "Germany", "Rating", "3", "City", "Darmstadt"),
		nodeInfoFromAttributes("Country", "Germany", "Rating", "7", "City", "Frankfurt"),
		nodeInfoFromAttributes("Country", "Russia", "Rating", "9", "City", "SPB"),
		nodeInfoFromAttributes("Country", "Russia", "Rating", "9", "City", "SPB"),
	}

	nm, err := NewNetmap(NodesFromInfo(nodes))
	require.NoError(t, err)
	c := NewContext(nm)
	require.NoError(t, c.processFilters(p))
	require.NoError(t, c.processSelectors(p))

	for _, s := range p.Selectors() {
		sel := c.Selections[s.Name()]
		s := c.Selectors[s.Name()]
		bucketCount, nodesInBucket := GetNodesCount(p, s)
		targ := fmt.Sprintf("selector '%s'", s.Name())
		require.Equal(t, bucketCount, len(sel), targ)
		for _, res := range sel {
			require.Equal(t, nodesInBucket, len(res), targ)
			for j := range res {
				require.True(t, c.applyFilter(s.Filter(), res[j]), targ)
			}
		}
	}

}

func TestPlacementPolicy_ProcessSelectorsHRW(t *testing.T) {
	p := newPlacementPolicy(1, nil,
		[]*Selector{
			newSelector("Main", "Country", ClauseDistinct, 3, "*"),
		}, nil)

	// bucket weight order: RU > DE > FR
	nodes := []NodeInfo{
		nodeInfoFromAttributes("Country", "Germany", PriceAttr, "2", CapacityAttr, "10000"),
		nodeInfoFromAttributes("Country", "Germany", PriceAttr, "4", CapacityAttr, "1"),
		nodeInfoFromAttributes("Country", "France", PriceAttr, "3", CapacityAttr, "10"),
		nodeInfoFromAttributes("Country", "Russia", PriceAttr, "2", CapacityAttr, "10000"),
		nodeInfoFromAttributes("Country", "Russia", PriceAttr, "1", CapacityAttr, "10000"),
		nodeInfoFromAttributes("Country", "Russia", CapacityAttr, "10000"),
		nodeInfoFromAttributes("Country", "France", PriceAttr, "100", CapacityAttr, "1"),
		nodeInfoFromAttributes("Country", "France", PriceAttr, "7", CapacityAttr, "10000"),
		nodeInfoFromAttributes("Country", "Russia", PriceAttr, "2", CapacityAttr, "1"),
	}
	nm, err := NewNetmap(NodesFromInfo(nodes))
	require.NoError(t, err)
	c := NewContext(nm)
	c.setPivot([]byte("containerID"))
	c.weightFunc = newWeightFunc(newMaxNorm(10000), newReverseMinNorm(1))
	c.aggregator = newMaxAgg
	require.NoError(t, c.processFilters(p))
	require.NoError(t, c.processSelectors(p))

	cnt := c.Selections["Main"]
	expected := []Nodes{
		{{Index: 4, Capacity: 10000, Price: 1}}, // best RU
		{{Index: 0, Capacity: 10000, Price: 2}}, // best DE
		{{Index: 7, Capacity: 10000, Price: 7}}, // best FR
	}
	require.Equal(t, len(expected), len(cnt))
	for i := range expected {
		require.Equal(t, len(expected[i]), len(cnt[i]))
		require.Equal(t, expected[i][0].Index, cnt[i][0].Index)
		require.Equal(t, expected[i][0].Capacity, cnt[i][0].Capacity)
		require.Equal(t, expected[i][0].Price, cnt[i][0].Price)
	}

	res, err := nm.GetPlacementVectors(containerNodes(cnt), []byte("objectID"))
	require.NoError(t, err)
	require.Equal(t, res, cnt)
}

func TestPlacementPolicy_ProcessSelectorsInvalid(t *testing.T) {
	testCases := []struct {
		name string
		p    *PlacementPolicy
		err  error
	}{
		{
			"MissingSelector",
			newPlacementPolicy(2, nil,
				[]*Selector{nil},
				[]*Filter{}),
			ErrMissingField,
		},
		{
			"InvalidFilterReference",
			newPlacementPolicy(1, nil,
				[]*Selector{newSelector("MyStore", "Country", ClauseDistinct, 1, "FromNL")},
				[]*Filter{newFilter("FromRU", "Country", "Russia", OpEQ)}),
			ErrFilterNotFound,
		},
		{
			"NotEnoughNodes (backup factor)",
			newPlacementPolicy(2, nil,
				[]*Selector{newSelector("MyStore", "Country", ClauseDistinct, 1, "FromRU")},
				[]*Filter{newFilter("FromRU", "Country", "Russia", OpEQ)}),
			ErrNotEnoughNodes,
		},
		{
			"NotEnoughNodes (buckets)",
			newPlacementPolicy(1, nil,
				[]*Selector{newSelector("MyStore", "Country", ClauseDistinct, 2, "FromRU")},
				[]*Filter{newFilter("FromRU", "Country", "Russia", OpEQ)}),
			ErrNotEnoughNodes,
		},
	}
	nodes := []NodeInfo{
		nodeInfoFromAttributes("Country", "Russia"),
		nodeInfoFromAttributes("Country", "Germany"),
		nodeInfoFromAttributes(),
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nm, err := NewNetmap(NodesFromInfo(nodes))
			require.NoError(t, err)
			c := NewContext(nm)
			require.NoError(t, c.processFilters(tc.p))

			err = c.processSelectors(tc.p)
			require.True(t, errors.Is(err, tc.err), "got: %v", err)
		})
	}
}

func testSelector() *Selector {
	s := new(Selector)
	s.SetName("name")
	s.SetCount(3)
	s.SetFilter("filter")
	s.SetAttribute("attribute")
	s.SetClause(ClauseDistinct)

	return s
}

func TestSelectorFromV2(t *testing.T) {
	sV2 := new(netmap.Selector)
	sV2.SetName("name")
	sV2.SetCount(3)
	sV2.SetClause(netmap.Distinct)
	sV2.SetAttribute("attribute")
	sV2.SetFilter("filter")

	s := NewSelectorFromV2(sV2)

	require.Equal(t, sV2, s.ToV2())
}

func TestSelector_Name(t *testing.T) {
	s := NewSelector()
	name := "some name"

	s.SetName(name)

	require.Equal(t, name, s.Name())
}

func TestSelector_Count(t *testing.T) {
	s := NewSelector()
	c := uint32(3)

	s.SetCount(c)

	require.Equal(t, c, s.Count())
}

func TestSelector_Clause(t *testing.T) {
	s := NewSelector()
	c := ClauseSame

	s.SetClause(c)

	require.Equal(t, c, s.Clause())
}

func TestSelector_Attribute(t *testing.T) {
	s := NewSelector()
	a := "some attribute"

	s.SetAttribute(a)

	require.Equal(t, a, s.Attribute())
}

func TestSelector_Filter(t *testing.T) {
	s := NewSelector()
	f := "some filter"

	s.SetFilter(f)

	require.Equal(t, f, s.Filter())
}
