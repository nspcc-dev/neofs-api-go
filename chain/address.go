package chain

import (
	"crypto/ecdsa"

	"github.com/nspcc-dev/neo-go/pkg/crypto/keys"
)

// WalletAddress implements NEO address.
type WalletAddress [AddressLength]byte

const (
	// AddressLength contains size of address,
	// 1 byte of address version + 20 bytes of ScriptHash + 4 bytes of checksum.
	AddressLength = 25
)

// KeyToAddress returns NEO address composed from public key.
func KeyToAddress(key *ecdsa.PublicKey) string {
	if key == nil {
		return ""
	}

	neoPublicKey := keys.PublicKey{
		X: key.X,
		Y: key.Y,
	}

	return neoPublicKey.Address()
}
