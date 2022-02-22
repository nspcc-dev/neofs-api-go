package signature

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	crypto "github.com/nspcc-dev/neofs-crypto"
)

type cfg struct {
	defaultScheme  refs.SignatureScheme
	restrictScheme refs.SignatureScheme
}

func defaultCfg() *cfg {
	return &cfg{
		defaultScheme:  refs.ECDSA_SHA512,
		restrictScheme: refs.UnspecifiedScheme,
	}
}

func verify(cfg *cfg, data []byte, sig *refs.Signature) error {
	scheme := sig.GetScheme()
	if scheme == refs.UnspecifiedScheme {
		scheme = cfg.defaultScheme
	}
	if cfg.restrictScheme != refs.UnspecifiedScheme && scheme != cfg.restrictScheme {
		return fmt.Errorf("%w: unexpected signature scheme", crypto.ErrInvalidSignature)
	}

	pub := crypto.UnmarshalPublicKey(sig.GetKey())
	switch scheme {
	case refs.ECDSA_SHA512:
		return crypto.Verify(pub, data, sig.GetSign())
	case refs.ECDSA_RFC6979_SHA256:
		return crypto.VerifyRFC6979(pub, data, sig.GetSign())
	default:
		return crypto.ErrInvalidSignature
	}
}

func sign(cfg *cfg, scheme refs.SignatureScheme, key *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	switch scheme {
	case refs.ECDSA_SHA512:
		return crypto.Sign(key, data)
	case refs.ECDSA_RFC6979_SHA256:
		return crypto.SignRFC6979(key, data)
	default:
		panic("unsupported scheme")
	}
}

func SignWithRFC6979() SignOption {
	return func(c *cfg) {
		c.defaultScheme = refs.ECDSA_RFC6979_SHA256
		c.restrictScheme = refs.ECDSA_RFC6979_SHA256
	}
}
