package netmap_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/stretchr/testify/require"
)

func TestNodeInfoJSON(t *testing.T) {
	exp := generateNodeInfo("public key", "/multi/addr", 2)

	t.Run("non empty", func(t *testing.T) {
		data, err := netmap.NodeInfoToJSON(exp)
		require.NoError(t, err)

		got, err := netmap.NodeInfoFromJSON(data)
		require.NoError(t, err)

		require.Equal(t, exp, got)
	})

	t.Run("empty", func(t *testing.T) {
		_, err := netmap.NodeInfoToJSON(nil)
		require.Error(t, err)

		_, err = netmap.NodeInfoFromJSON(nil)
		require.Error(t, err)
	})
}
