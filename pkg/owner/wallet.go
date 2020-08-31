package owner

import (
	"crypto/ecdsa"

	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neo-go/pkg/crypto/keys"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/pkg/errors"
)

// NEO3Wallet represents NEO3 wallet address.
type NEO3Wallet [25]byte

// NEO3WalletFromPublicKey converts public key to NEO3 wallet address.
func NEO3WalletFromPublicKey(key *ecdsa.PublicKey) (*NEO3Wallet, error) {
	if key == nil {
		return nil, crypto.ErrEmptyPublicKey
	}

	neoPublicKey := keys.PublicKey{
		X: key.X,
		Y: key.Y,
	}

	d, err := base58.Decode(neoPublicKey.Address())
	if err != nil {
		return nil, errors.Wrap(err, "can't decode neo3 address from key")
	}

	w := new(NEO3Wallet)

	copy(w.Bytes(), d)

	return w, nil
}

func (w *NEO3Wallet) String() string {
	if w != nil {
		return base58.Encode(w[:])
	}

	return ""
}

// Bytes returns slice of NEO3 wallet address bytes.
func (w *NEO3Wallet) Bytes() []byte {
	if w != nil {
		return w[:]
	}

	return nil
}
