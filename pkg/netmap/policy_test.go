package netmap

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlacementPolicy_CBFWithEmptySelector(t *testing.T) {
	nodes := []NodeInfo{
		nodeInfoFromAttributes("ID", "1", "Attr", "Same"),
		nodeInfoFromAttributes("ID", "2", "Attr", "Same"),
		nodeInfoFromAttributes("ID", "3", "Attr", "Same"),
		nodeInfoFromAttributes("ID", "4", "Attr", "Same"),
	}

	p1 := newPlacementPolicy(0,
		[]*Replica{newReplica(2, "")},
		nil, // selectors
		nil, // filters
	)

	p2 := newPlacementPolicy(3,
		[]*Replica{newReplica(2, "")},
		nil, // selectors
		nil, // filters
	)

	p3 := newPlacementPolicy(3,
		[]*Replica{newReplica(2, "X")},
		[]*Selector{newSelector("X", "", ClauseDistinct, 2, "*")},
		nil, // filters
	)

	p4 := newPlacementPolicy(3,
		[]*Replica{newReplica(2, "X")},
		[]*Selector{newSelector("X", "Attr", ClauseSame, 2, "*")},
		nil, // filters
	)

	nm, err := NewNetmap(NodesFromInfo(nodes))
	require.NoError(t, err)

	v, err := nm.GetContainerNodes(p1, nil)
	require.NoError(t, err)
	assert.Len(t, v.Flatten(), 4)

	v, err = nm.GetContainerNodes(p2, nil)
	require.NoError(t, err)
	assert.Len(t, v.Flatten(), 4)

	v, err = nm.GetContainerNodes(p3, nil)
	require.NoError(t, err)
	assert.Len(t, v.Flatten(), 4)

	v, err = nm.GetContainerNodes(p4, nil)
	require.NoError(t, err)
	assert.Len(t, v.Flatten(), 4)
}

func TestPlacementPolicyFromV2(t *testing.T) {
	pV2 := new(netmap.PlacementPolicy)

	pV2.SetReplicas([]*netmap.Replica{
		testReplica().ToV2(),
		testReplica().ToV2(),
	})

	pV2.SetContainerBackupFactor(3)

	pV2.SetSelectors([]*netmap.Selector{
		testSelector().ToV2(),
		testSelector().ToV2(),
	})

	pV2.SetFilters([]*netmap.Filter{
		testFilter().ToV2(),
		testFilter().ToV2(),
	})

	p := NewPlacementPolicyFromV2(pV2)

	require.Equal(t, pV2, p.ToV2())
}

func TestPlacementPolicy_Replicas(t *testing.T) {
	p := NewPlacementPolicy()
	rs := []*Replica{testReplica(), testReplica()}

	p.SetReplicas(rs...)

	require.Equal(t, rs, p.Replicas())
}

func TestPlacementPolicy_ContainerBackupFactor(t *testing.T) {
	p := NewPlacementPolicy()
	f := uint32(3)

	p.SetContainerBackupFactor(f)

	require.Equal(t, f, p.ContainerBackupFactor())
}

func TestPlacementPolicy_Selectors(t *testing.T) {
	p := NewPlacementPolicy()
	ss := []*Selector{testSelector(), testSelector()}

	p.SetSelectors(ss...)

	require.Equal(t, ss, p.Selectors())
}

func TestPlacementPolicy_Filters(t *testing.T) {
	p := NewPlacementPolicy()
	fs := []*Filter{testFilter(), testFilter()}

	p.SetFilters(fs...)

	require.Equal(t, fs, p.Filters())
}

func TestPlacementPolicyEncoding(t *testing.T) {
	p := newPlacementPolicy(3, nil, nil, nil)

	t.Run("binary", func(t *testing.T) {
		data, err := p.Marshal()
		require.NoError(t, err)

		p2 := NewPlacementPolicy()
		require.NoError(t, p2.Unmarshal(data))

		require.Equal(t, p, p2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := p.MarshalJSON()
		require.NoError(t, err)

		p2 := NewPlacementPolicy()
		require.NoError(t, p2.UnmarshalJSON(data))

		require.Equal(t, p, p2)
	})
}

func TestNewPlacementPolicy(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *PlacementPolicy

		require.Nil(t, x.ToV2())
	})

	t.Run("default values", func(t *testing.T) {
		pp := NewPlacementPolicy()

		// check initial values
		require.Nil(t, pp.Replicas())
		require.Nil(t, pp.Filters())
		require.Nil(t, pp.Selectors())
		require.Zero(t, pp.ContainerBackupFactor())

		// convert to v2 message
		ppV2 := pp.ToV2()

		require.Nil(t, ppV2.GetReplicas())
		require.Nil(t, ppV2.GetFilters())
		require.Nil(t, ppV2.GetSelectors())
		require.Zero(t, ppV2.GetContainerBackupFactor())
	})
}

func TestNewPlacementPolicyFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *netmap.PlacementPolicy

		require.Nil(t, NewPlacementPolicyFromV2(x))
	})
}
