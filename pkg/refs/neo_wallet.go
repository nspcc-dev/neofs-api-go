package refs

import (
	"crypto/ecdsa"

	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neo-go/pkg/crypto/keys"
	"github.com/pkg/errors"
)

type (
	NEO3Wallet [25]byte
)

func NEO3WalletFromPublicKey(key *ecdsa.PublicKey) (owner NEO3Wallet, err error) {
	if key == nil {
		return owner, errors.New("nil public key")
	}

	neoPublicKey := keys.PublicKey{
		X: key.X,
		Y: key.Y,
	}

	d, err := base58.Decode(neoPublicKey.Address())
	if err != nil {
		return owner, errors.Wrap(err, "can't decode neo3 address from key")
	}

	copy(owner[:], d)

	return owner, nil
}

func (w NEO3Wallet) String() string {
	return base58.Encode(w[:])
}
