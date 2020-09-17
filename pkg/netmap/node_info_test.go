package netmap

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

func TestNode_NetworkAddress(t *testing.T) {
	addr := "127.0.0.1:8080"

	nV2 := new(netmap.NodeInfo)
	nV2.SetAddress(addr)

	n := Node{
		InfoV2: nV2,
	}

	require.Equal(t, addr, n.NetworkAddress())
}
