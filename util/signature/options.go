package signature

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	crypto "github.com/nspcc-dev/neofs-crypto"
)

type cfg struct {
	schemeFixed bool
	scheme      refs.SignatureScheme
}

func defaultCfg() *cfg {
	return new(cfg)
}

func verify(cfg *cfg, data []byte, sig *refs.Signature) error {
	if !cfg.schemeFixed {
		cfg.scheme = sig.GetScheme()
	}

	pub := crypto.UnmarshalPublicKey(sig.GetKey())

	switch cfg.scheme {
	case refs.ECDSA_SHA512:
		return crypto.Verify(pub, data, sig.GetSign())
	case refs.ECDSA_RFC6979_SHA256:
		return crypto.VerifyRFC6979(pub, data, sig.GetSign())
	default:
		return fmt.Errorf("unsupported signature scheme %s", cfg.scheme)
	}
}

func sign(cfg *cfg, key *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	switch cfg.scheme {
	case refs.ECDSA_SHA512:
		return crypto.Sign(key, data)
	case refs.ECDSA_RFC6979_SHA256:
		return crypto.SignRFC6979(key, data)
	default:
		panic(fmt.Sprintf("unsupported scheme %s", cfg.scheme))
	}
}

func SignWithRFC6979() SignOption {
	return func(c *cfg) {
		c.schemeFixed = true
		c.scheme = refs.ECDSA_RFC6979_SHA256
	}
}
