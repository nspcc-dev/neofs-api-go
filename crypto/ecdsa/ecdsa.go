package neofsecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"math/big"
	"sync"

	neofscrypto "github.com/nspcc-dev/neofs-api-go/crypto"
)

type signer ecdsa.PrivateKey

var (
	onceCurve sync.Once
	curve     elliptic.Curve
)

// Signer provides neofscrypto.Signer interface of the ecdsa.PrivateKey.
//
// Public key is based on ecdsa.PublicKey.
//
// Signature is calculated via ecdsa.Sign and elliptic.Marshal
// functions over the result of SHA512 cryptographic hash function
// on data.
//
// Uses elliptic.P256() elliptic curve.
func Signer(key ecdsa.PrivateKey) neofscrypto.Signer {
	return (signer)(key)
}

func (x signer) Sign(data []byte) ([]byte, error) {
	onceCurve.Do(func() { curve = elliptic.P256() })

	h := sha512.Sum512(data)

	r, s, err := ecdsa.Sign(rand.Reader, (*ecdsa.PrivateKey)(&x), h[:])
	if err != nil {
		return nil, err
	}

	return elliptic.Marshal(curve, r, s), nil
}

func (x signer) Public() neofscrypto.PublicKey {
	return Public((ecdsa.PublicKey)(x.PublicKey))
}

type publicKey ecdsa.PublicKey

// Public provides neofscrypto.PublicKey interface of the ecdsa.PublicKey,
//
// Verifies signature via ecdsa.Verify function.
//
// Marshals key to a binary format via elliptic.Marshal.
//
// Uses elliptic.P256() elliptic curve.
func Public(key ecdsa.PublicKey) neofscrypto.PublicKey {
	return (*publicKey)(&key)
}

// PublicBlank provides neofscrypto.PublicKey interface
// of the blank ecdsa.PublicKey.
//
// Commonly used to unmarshal the binary representation.
func PublicBlank() neofscrypto.PublicKey {
	return Public(ecdsa.PublicKey{})
}

func (x publicKey) Verify(data, signature []byte) bool {
	onceCurve.Do(func() { curve = elliptic.P256() })

	var r, s *big.Int

	{
		// code is copy-pasted from elliptic.Unmarshal
		curve := elliptic.P256()
		curvePrm := curve.Params()

		byteLen := (curvePrm.BitSize + 7) / 8
		if len(signature) != 1+2*byteLen {
			return false
		}

		if signature[0] != 4 { // uncompressed form
			return false
		}

		r = new(big.Int).SetBytes(signature[1 : 1+byteLen])
		s = new(big.Int).SetBytes(signature[1+byteLen:])

		if r.Cmp(curvePrm.P) >= 0 || s.Cmp(curvePrm.P) >= 0 {
			return false
		}

		// removed statement: r and s are not on curve even after ecdsa.Sign
		// if !curve.IsOnCurve(r, s) {
		// 	return false
		// }
	}

	h := sha512.Sum512(data)

	return ecdsa.Verify((*ecdsa.PublicKey)(&x), h[:], r, s)
}

func (x publicKey) MarshalBinary() (data []byte, err error) {
	onceCurve.Do(func() { curve = elliptic.P256() })

	k := (ecdsa.PublicKey)(x)

	return elliptic.MarshalCompressed(curve, k.X, k.Y), nil
}

func (x *publicKey) UnmarshalBinary(data []byte) error {
	return UnmarshalPublicKey((*ecdsa.PublicKey)(x), data)
}

// UnmarshalPublicKey unmarshals ecdsa.PublicKey from a binary representation.
//
// Key must not be nil.
func UnmarshalPublicKey(key *ecdsa.PublicKey, data []byte) error {
	onceCurve.Do(func() { curve = elliptic.P256() })

	a, b := elliptic.UnmarshalCompressed(curve, data)
	if a == nil {
		return errors.New("point not on the curve")
	}

	key.Curve = curve
	key.X = a
	key.Y = b

	return nil
}
