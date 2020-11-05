package netmap

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

func testReplica() *Replica {
	r := new(Replica)
	r.SetCount(3)
	r.SetSelector("selector")

	return r
}

func TestReplicaFromV2(t *testing.T) {
	rV2 := new(netmap.Replica)
	rV2.SetCount(3)
	rV2.SetSelector("selector")

	r := NewReplicaFromV2(rV2)

	require.Equal(t, rV2, r.ToV2())
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
