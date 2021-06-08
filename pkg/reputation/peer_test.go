package reputation_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/reputation"
	reputationtest "github.com/nspcc-dev/neofs-api-go/pkg/reputation/test"
	reputationV2 "github.com/nspcc-dev/neofs-api-go/v2/reputation"
	"github.com/stretchr/testify/require"
)

func TestPeerID_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *reputation.PeerID

		require.Nil(t, x.ToV2())
	})

	t.Run("nil", func(t *testing.T) {
		peerID := reputationtest.GeneratePeerID()

		require.Equal(t, peerID, reputation.PeerIDFromV2(peerID.ToV2()))
	})
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

func TestPeerIDFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *reputationV2.PeerID

		require.Nil(t, reputation.PeerIDFromV2(x))
	})
}
