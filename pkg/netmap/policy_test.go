package netmap

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

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
