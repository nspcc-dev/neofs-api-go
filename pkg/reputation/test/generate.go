package reputationtest

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/reputation"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func GeneratePeerID() *reputation.PeerID {
	v := reputation.NewPeerID()

	key := [crypto.PublicKeyCompressedSize]byte{}
	copy(key[:], crypto.MarshalPublicKey(&test.DecodeKey(-1).PublicKey))

	v.SetPublicKey(key)

	return v
}

func GenerateTrust() *reputation.Trust {
	v := reputation.NewTrust()
	v.SetPeer(GeneratePeerID())
	v.SetValue(1.5)

	return v
}

func GenerateGlobalTrust() *reputation.GlobalTrust {
	v := reputation.NewGlobalTrust()
	v.SetManager(GeneratePeerID())
	v.SetTrust(GenerateTrust())

	return v
}

func GenerateSignedGlobalTrust(t testing.TB) *reputation.GlobalTrust {
	gt := GenerateGlobalTrust()

	require.NoError(t, gt.Sign(test.DecodeKey(0)))

	return gt
}
