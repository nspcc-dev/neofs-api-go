package service

import (
	"crypto/ecdsa"

	crypto "github.com/nspcc-dev/neofs-crypto"
)

type keySign struct {
	key  *ecdsa.PublicKey
	sign []byte
}

// GetSignature is a sign field getter.
func (s keySign) GetSignature() []byte {
	return s.sign
}

// GetPublicKey is a key field getter,
func (s keySign) GetPublicKey() *ecdsa.PublicKey {
	return s.key
}

// Unites passed key with signature and returns SignKeyPair interface.
func newSignatureKeyPair(key *ecdsa.PublicKey, sign []byte) SignKeyPair {
	return &keySign{
		key:  key,
		sign: sign,
	}
}

// Returns data from DataSignatureAccumulator for signature creation/verification.
//
// If passed DataSignatureAccumulator provides a SignedDataReader interface, data for signature is obtained
// using this interface for optimization. In this case, it is understood that reading into the slice D
// that the method DataForSignature returns does not change D.
func dataForSignature(src SignedDataSource) ([]byte, error) {
	if src == nil {
		return nil, ErrNilSignedDataSource
	}

	r, ok := src.(SignedDataReader)
	if !ok {
		return src.SignedData()
	}

	buf := bytesPool.Get().([]byte)
	defer func() {
		bytesPool.Put(buf)
	}()

	if size := r.SignedDataSize(); size <= cap(buf) {
		buf = buf[:size]
	} else {
		buf = make([]byte, size)
	}

	n, err := r.ReadSignedData(buf)
	if err != nil {
		return nil, err
	}

	return buf[:n], nil

}

// DataSignature returns the signature of data obtained using the private key.
//
// If passed data container is nil, ErrNilSignedDataSource returns.
// If passed private key is nil, crypto.ErrEmptyPrivateKey returns.
// If the data container or the signature function returns an error, it is returned directly.
func DataSignature(key *ecdsa.PrivateKey, src SignedDataSource) ([]byte, error) {
	if key == nil {
		return nil, crypto.ErrEmptyPrivateKey
	}

	data, err := dataForSignature(src)
	if err != nil {
		return nil, err
	}

	return crypto.Sign(key, data)
}

// AddSignatureWithKey calculates the data signature and adds it to accumulator with public key.
//
// Returns signing errors only.
func AddSignatureWithKey(v SignatureKeyAccumulator, key *ecdsa.PrivateKey) error {
	sign, err := DataSignature(key, v)
	if err != nil {
		return err
	}

	v.AddSignKey(sign, &key.PublicKey)

	return nil
}

// Checks passed key-signature pairs for data from the passed container.
//
// If passed key-signatures pair set is empty, nil returns immediately.
func verifySignatures(src SignedDataSource, items ...SignKeyPair) error {
	if len(items) <= 0 {
		return nil
	}

	data, err := dataForSignature(src)
	if err != nil {
		return err
	}

	for _, signKey := range items {
		if err := crypto.Verify(
			signKey.GetPublicKey(),
			data,
			signKey.GetSignature(),
		); err != nil {
			return err
		}
	}

	return nil
}

// VerifySignatures checks passed key-signature pairs for data from the passed container.
//
// If passed data source is nil, ErrNilSignedDataSource returns.
// If check data is not ready, corresponding error returns.
// If at least one of the pairs is invalid, an error returns.
func VerifySignatures(src SignedDataSource, items ...SignKeyPair) error {
	return verifySignatures(src, items...)
}

// VerifyAccumulatedSignatures checks if accumulated key-signature pairs are valid.
//
// Behaves like VerifySignatures.
// If passed key-signature source is empty, ErrNilSignatureKeySource returns.
func VerifyAccumulatedSignatures(src SignatureKeySource) error {
	if src == nil {
		return ErrNilSignatureKeySource
	}

	return verifySignatures(src, src.GetSignKeyPairs()...)
}

// VerifySignatureWithKey checks data signature from the passed container with passed key.
//
// If passed data with signature is nil, ErrEmptyDataWithSignature returns.
// If passed key is nil, crypto.ErrEmptyPublicKey returns.
// A non-nil error returns if and only if the signature does not pass verification.
func VerifySignatureWithKey(src DataWithSignature, key *ecdsa.PublicKey) error {
	if src == nil {
		return ErrEmptyDataWithSignature
	} else if key == nil {
		return crypto.ErrEmptyPublicKey
	}

	return verifySignatures(
		src,
		newSignatureKeyPair(
			key,
			src.GetSignature(),
		),
	)
}
