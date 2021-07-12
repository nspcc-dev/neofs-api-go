package neofsrfc6979

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"errors"
	"math/big"
	"sync"

	neofscrypto "github.com/nspcc-dev/neofs-api-go/crypto"
	"github.com/nspcc-dev/rfc6979"
)

type rfc6979Signer ecdsa.PrivateKey

// Signer provides neofscrypto.Signer interface of the ecdsa.PrivateKey.
//
// Public key is based on ecdsa.PublicKey.
//
// Signature is calculated via rfc6979.SignECDSA and elliptic.Marshal
// functions over the result of SHA256 cryptographic hash function
// on data.
//
// Uses elliptic.P256() elliptic curve.
func Signer(key ecdsa.PrivateKey) neofscrypto.Signer {
	return (rfc6979Signer)(key)
}

var (
	onceCurve sync.Once
	curve     elliptic.Curve
)

func initCurve() {
	curve = elliptic.P256()
}

const rfc6979SigLen = 64

func (x rfc6979Signer) Sign(data []byte) ([]byte, error) {
	onceCurve.Do(initCurve)

	h := sha256.Sum256(data)

	r, s := rfc6979.SignECDSA((*ecdsa.PrivateKey)(&x), h[:], sha256.New)

	mid := rfc6979SigLen / 2
	sig := make([]byte, rfc6979SigLen)

	chunk := r.Bytes()
	copy(sig[mid-len(chunk):], chunk)

	chunk = s.Bytes()
	copy(sig[rfc6979SigLen-len(chunk):], chunk)

	return sig, nil
}

func (x rfc6979Signer) Public() neofscrypto.PublicKey {
	return (*rfc9679PublicKey)(&x.PublicKey)
}

type rfc9679PublicKey ecdsa.PublicKey

// Public provides neofscrypto.PublicKey interface of the ecdsa.PublicKey,
//
// Verifies signature via ecdsa.Verify function.
//
// Marshals key to a binary format via elliptic.Marshal.
//
// Uses elliptic.P256() elliptic curve.
func Public(key ecdsa.PublicKey) neofscrypto.PublicKey {
	return (*rfc9679PublicKey)(&key)
}

// PublicBlank provides neofscrypto.PublicKey interface
// of the blank ecdsa.PublicKey.
//
// Commonly used to unmarshal the binary representation.
func PublicBlank() neofscrypto.PublicKey {
	return Public(ecdsa.PublicKey{})
}

func (x *rfc9679PublicKey) Verify(data, signature []byte) bool {
	mid := rfc6979SigLen / 2
	if len(signature) < mid {
		return false
	}

	r, s := new(big.Int).SetBytes(signature[:mid]), new(big.Int).SetBytes(signature[mid:])

	h := sha256.Sum256(data)

	return ecdsa.Verify((*ecdsa.PublicKey)(x), h[:], r, s)
}

func (x rfc9679PublicKey) MarshalBinary() (data []byte, err error) {
	onceCurve.Do(initCurve)

	k := (ecdsa.PublicKey)(x)

	return elliptic.MarshalCompressed(curve, k.X, k.Y), nil
}

func (x *rfc9679PublicKey) UnmarshalBinary(data []byte) error {
	onceCurve.Do(initCurve)

	a, b := elliptic.UnmarshalCompressed(curve, data)
	if a == nil {
		return errors.New("point on the curve")
	}

	k := (*ecdsa.PublicKey)(x)
	k.Curve = curve
	k.X = a
	k.Y = b

	return nil
}
