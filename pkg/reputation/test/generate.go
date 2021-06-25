package reputationtest

import (
	"testing"

	neofsecdsatest "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/reputation"
	"github.com/stretchr/testify/require"
)

func GeneratePeerID() *reputation.PeerID {
	v := reputation.NewPeerID()

	v.SetPublicKey(neofsecdsatest.PublicBytes())

	return v
}

func GenerateTrust() *reputation.Trust {
	v := reputation.NewTrust()
	v.SetPeer(GeneratePeerID())
	v.SetValue(1.5)

	return v
}

func GeneratePeerToPeerTrust() *reputation.PeerToPeerTrust {
	v := reputation.NewPeerToPeerTrust()
	v.SetTrustingPeer(GeneratePeerID())
	v.SetTrust(GenerateTrust())

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

	require.NoError(t, gt.SignECDSA(neofsecdsatest.Key()))

	return gt
}
