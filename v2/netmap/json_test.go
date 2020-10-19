package netmap_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

func TestNodeInfoJSON(t *testing.T) {
	exp := generateNodeInfo("public key", "/multi/addr", 2)

	t.Run("non empty", func(t *testing.T) {
		data := netmap.NodeInfoToJSON(exp)
		require.NotNil(t, data)

		got := netmap.NodeInfoFromJSON(data)
		require.NotNil(t, got)

		require.Equal(t, exp, got)
	})
}
