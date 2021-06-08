package netmap

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	testv2 "github.com/nspcc-dev/neofs-api-go/v2/netmap/test"
	"github.com/stretchr/testify/require"
)

func testReplica() *Replica {
	r := new(Replica)
	r.SetCount(3)
	r.SetSelector("selector")

	return r
}

func TestReplicaFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *netmap.Replica

		require.Nil(t, NewReplicaFromV2(x))
	})

	t.Run("from non-nil", func(t *testing.T) {
		rV2 := testv2.GenerateReplica(false)

		r := NewReplicaFromV2(rV2)

		require.Equal(t, rV2, r.ToV2())
	})
}

func TestReplica_Count(t *testing.T) {
	r := NewReplica()
	c := uint32(3)

	r.SetCount(c)

	require.Equal(t, c, r.Count())
}

func TestReplica_Selector(t *testing.T) {
	r := NewReplica()
	s := "some selector"

	r.SetSelector(s)

	require.Equal(t, s, r.Selector())
}

func TestReplicaEncoding(t *testing.T) {
	r := newReplica(3, "selector")

	t.Run("binary", func(t *testing.T) {
		data, err := r.Marshal()
		require.NoError(t, err)

		r2 := NewReplica()
		require.NoError(t, r2.Unmarshal(data))

		require.Equal(t, r, r2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := r.MarshalJSON()
		require.NoError(t, err)

		r2 := NewReplica()
		require.NoError(t, r2.UnmarshalJSON(data))

		require.Equal(t, r, r2)
	})
}

func TestReplica_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *Replica

		require.Nil(t, x.ToV2())
	})
}

func TestNewReplica(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		r := NewReplica()

		// check initial values
		require.Zero(t, r.Count())
		require.Empty(t, r.Selector())

		// convert to v2 message
		rV2 := r.ToV2()

		require.Zero(t, rV2.GetCount())
		require.Empty(t, rV2.GetSelector())
	})
}
