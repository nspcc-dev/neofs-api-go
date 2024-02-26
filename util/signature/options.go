package signature

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/util/signature/walletconnect"
	"github.com/nspcc-dev/rfc6979"
)

const sigIntLen = 32

type cfg struct {
	schemeFixed bool
	scheme      refs.SignatureScheme
	buffer      []byte
}

func defaultCfg() *cfg {
	return new(cfg)
}

func verify(cfg *cfg, data []byte, sig *refs.Signature) error {
	if !cfg.schemeFixed {
		cfg.scheme = sig.GetScheme()
	}

	var pub ecdsa.PublicKey
	pub.Curve = elliptic.P256()
	pub.X, pub.Y = elliptic.UnmarshalCompressed(pub.Curve, sig.GetKey())
	if pub.X == nil {
		return errors.New("invalid public key")
	}

	var sigb = sig.GetSign()

	if cfg.scheme == refs.ECDSA_RFC6979_SHA256_WALLET_CONNECT {
		buf := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
		base64.StdEncoding.Encode(buf, data)
		if !walletconnect.Verify(&pub, buf, sigb) {
			return errors.New("invalid signature")
		}
		return nil
	}
	var (
		r, s    big.Int
		res     bool
		prefLen int
	)

	if cfg.scheme == refs.ECDSA_SHA512 {
		prefLen = 1
	}
	if len(sigb) != prefLen+2*sigIntLen {
		return errors.New("invalid signature length")
	}

	r.SetBytes(sigb[prefLen : prefLen+sigIntLen])
	s.SetBytes(sigb[prefLen+sigIntLen:])

	switch cfg.scheme {
	case refs.ECDSA_SHA512:
		var h = sha512.Sum512(data)

		if sigb[0] != 0x04 {
			return errors.New("invalid signature prefix") // Legacy prefix for SHA512 signature.
		}
		res = ecdsa.Verify(&pub, h[:], &r, &s)
	case refs.ECDSA_RFC6979_SHA256:
		var h = sha256.Sum256(data)
		res = ecdsa.Verify(&pub, h[:], &r, &s)
	default:
		return fmt.Errorf("unsupported signature scheme %s", cfg.scheme)
	}
	if !res {
		return errors.New("invalid signature")
	}
	return nil
}

func sign(cfg *cfg, key *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	if cfg.scheme == refs.ECDSA_RFC6979_SHA256_WALLET_CONNECT {
		buf := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
		base64.StdEncoding.Encode(buf, data)
		return walletconnect.Sign(key, buf)
	}

	var (
		r, s    *big.Int
		err     error
		prefLen int
	)
	switch cfg.scheme {
	case refs.ECDSA_SHA512:
		var h = sha512.Sum512(data)

		r, s, err = ecdsa.Sign(rand.Reader, key, h[:])
		prefLen = 1
	case refs.ECDSA_RFC6979_SHA256:
		var h = sha256.Sum256(data)
		r, s = rfc6979.SignECDSA(key, h[:], sha256.New)
	default:
		panic(fmt.Sprintf("unsupported scheme %s", cfg.scheme))
	}
	if err != nil {
		return nil, err
	}
	var sig = make([]byte, prefLen+2*sigIntLen)
	if prefLen != 0 {
		sig[0] = 0x04 // Legacy prefix.
	}
	r.FillBytes(sig[prefLen : prefLen+sigIntLen])
	s.FillBytes(sig[prefLen+sigIntLen:])
	return sig, nil
}

func SignWithRFC6979() SignOption {
	return func(c *cfg) {
		c.schemeFixed = true
		c.scheme = refs.ECDSA_RFC6979_SHA256
	}
}

// WithBuffer allows providing pre-allocated buffer for signature verification.
func WithBuffer(buf []byte) SignOption {
	return func(c *cfg) {
		c.buffer = buf
	}
}

func SignWithWalletConnect() SignOption {
	return func(c *cfg) {
		c.scheme = refs.ECDSA_RFC6979_SHA256_WALLET_CONNECT
	}
}
