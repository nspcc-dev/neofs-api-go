package reputation_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/reputation"
	"github.com/stretchr/testify/require"
)

func TestPeerID(t *testing.T) {
	peerID := reputation.NewPeerID()

	data := []byte{1, 2, 3}
	peerID.SetBytes(data)
	require.Equal(t, data, peerID.Bytes())

	require.Equal(t, peerID, reputation.PeerIDFromV2(peerID.ToV2()))
}
