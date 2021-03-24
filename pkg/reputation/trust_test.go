package reputation_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/reputation"
	reputationtest "github.com/nspcc-dev/neofs-api-go/pkg/reputation/test"
	"github.com/stretchr/testify/require"
)

func TestTrust(t *testing.T) {
	trust := reputation.NewTrust()

	id := reputationtest.GeneratePeerID()
	trust.SetPeer(id)
	require.Equal(t, id, trust.Peer())

	val := 1.5
	trust.SetValue(val)
	require.Equal(t, val, trust.Value())

	t.Run("binary encoding", func(t *testing.T) {
		trust := reputationtest.GenerateTrust()
		data, err := trust.Marshal()
		require.NoError(t, err)

		trust2 := reputation.NewTrust()
		require.NoError(t, trust2.Unmarshal(data))
		require.Equal(t, trust, trust2)
	})

	t.Run("JSON encoding", func(t *testing.T) {
		trust := reputationtest.GenerateTrust()
		data, err := trust.MarshalJSON()
		require.NoError(t, err)

		trust2 := reputation.NewTrust()
		require.NoError(t, trust2.UnmarshalJSON(data))
		require.Equal(t, trust, trust2)
	})
}
