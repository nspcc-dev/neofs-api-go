package signature

import (
	"crypto/ecdsa"

	crypto "github.com/nspcc-dev/neofs-crypto"
)

type cfg struct {
	signFunc   func(key *ecdsa.PrivateKey, msg []byte) ([]byte, error)
	verifyFunc func(key *ecdsa.PublicKey, msg []byte, sig []byte) error
}

func defaultCfg() *cfg {
	return &cfg{
		signFunc:   crypto.Sign,
		verifyFunc: crypto.Verify,
	}
}

func SignWithRFC6979() SignOption {
	return func(c *cfg) {
		c.signFunc = crypto.SignRFC6979
		c.verifyFunc = crypto.VerifyRFC6979
	}
}
