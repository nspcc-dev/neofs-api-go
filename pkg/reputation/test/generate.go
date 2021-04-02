package reputationtest

import (
	"github.com/nspcc-dev/neofs-api-go/pkg/reputation"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-crypto/test"
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
