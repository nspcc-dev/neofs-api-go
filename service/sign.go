package service

import (
	"crypto/ecdsa"

	crypto "github.com/nspcc-dev/neofs-crypto"
)

// Returns data from DataSignatureAccumulator for signature creation/verification.
//
// If passed DataSignatureAccumulator provides a SignedDataReader interface, data for signature is obtained
// using this interface for optimization. In this case, it is understood that reading into the slice D
// that the method DataForSignature returns does not change D.
func dataForSignature(src SignedDataSource) ([]byte, error) {
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
func DataSignature(src SignedDataSource, key *ecdsa.PrivateKey) ([]byte, error) {
	if src == nil {
		return nil, ErrNilSignedDataSource
	} else if key == nil {
		return nil, crypto.ErrEmptyPrivateKey
	}

	data, err := dataForSignature(src)
	if err != nil {
		return nil, err
	}

	return crypto.Sign(key, data)
}
