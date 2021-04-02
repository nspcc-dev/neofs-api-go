package reputation_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/reputation"
	reputationtest "github.com/nspcc-dev/neofs-api-go/pkg/reputation/test"
	"github.com/stretchr/testify/require"
)

func TestPeerID_ToV2(t *testing.T) {
	peerID := reputationtest.GeneratePeerID()

	require.Equal(t, peerID, reputation.PeerIDFromV2(peerID.ToV2()))
}

func TestPeerID_String(t *testing.T) {
	id := reputationtest.GeneratePeerID()

	strID := id.String()

	id2 := reputation.NewPeerID()

	err := id2.Parse(strID)
	require.NoError(t, err)

	require.Equal(t, id, id2)
}

func TestPeerIDEncoding(t *testing.T) {
	id := reputationtest.GeneratePeerID()

	t.Run("binary", func(t *testing.T) {
		data, err := id.Marshal()
		require.NoError(t, err)

		id2 := reputation.NewPeerID()
		require.NoError(t, id2.Unmarshal(data))

		require.Equal(t, id, id2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := id.MarshalJSON()
		require.NoError(t, err)

		id2 := reputation.NewPeerID()
		require.NoError(t, id2.UnmarshalJSON(data))

		require.Equal(t, id, id2)
	})
}
