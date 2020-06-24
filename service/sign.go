package service

import (
	"crypto/ecdsa"
	"io"
	"sync"

	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/pkg/errors"
)

type keySign struct {
	key  *ecdsa.PublicKey
	sign []byte
}

type signSourceGroup struct {
	SignKeyPairSource
	SignKeyPairAccumulator

	sources []SignedDataSource
}

type signReadersGroup struct {
	SignKeyPairSource
	SignKeyPairAccumulator

	readers []SignedDataReader
}

var bytesPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 5<<20)
	},
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
//
// If returned length of data is negative, ErrNegativeLength returns.
func dataForSignature(src SignedDataSource) ([]byte, error) {
	if src == nil {
		return nil, ErrNilSignedDataSource
	}

	r, ok := src.(SignedDataReader)
	if !ok {
		return src.SignedData()
	}

	buf := bytesPool.Get().([]byte)

	if size := r.SignedDataSize(); size < 0 {
		return nil, ErrNegativeLength
	} else if size <= cap(buf) {
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
	defer bytesPool.Put(data)

	return crypto.Sign(key, data)
}

// AddSignatureWithKey calculates the data signature and adds it to accumulator with public key.
//
// Any change of data provoke signature breakdown.
//
// Returns signing errors only.
func AddSignatureWithKey(key *ecdsa.PrivateKey, v DataWithSignKeyAccumulator) error {
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
	defer bytesPool.Put(data)

	for i := range items {
		if i > 0 {
			// add previous key bytes to the signed message

			signKeyDataSrc := SignKeyPairsSignedData(items[i-1])

			signKeyData, err := signKeyDataSrc.SignedData()
			if err != nil {
				return errors.Wrapf(err, "could not get signed data of key-signature #%d", i)
			}

			data = append(data, signKeyData...)
		}

		if err := crypto.Verify(
			items[i].GetPublicKey(),
			data,
			items[i].GetSignature(),
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
func VerifyAccumulatedSignatures(src DataWithSignKeySource) error {
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
func VerifySignatureWithKey(key *ecdsa.PublicKey, src DataWithSignature) error {
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

// SignRequestData calculates request data signature and adds it to accumulator.
//
// Any change of request data provoke signature breakdown.
//
// If passed private key is nil, crypto.ErrEmptyPrivateKey returns.
// If passed RequestSignedData is nil, ErrNilRequestSignedData returns.
func SignRequestData(key *ecdsa.PrivateKey, src RequestSignedData) error {
	if src == nil {
		return ErrNilRequestSignedData
	}

	sigSrc, err := GroupSignedPayloads(
		src,
		src,
		NewSignedSessionToken(
			src.GetSessionToken(),
		),
		NewSignedBearerToken(
			src.GetBearerToken(),
		),
		ExtendedHeadersSignedData(src),
		SignKeyPairsSignedData(src.GetSignKeyPairs()...),
	)
	if err != nil {
		return err
	}

	return AddSignatureWithKey(key, sigSrc)
}

// VerifyRequestData checks if accumulated key-signature pairs of data with token are valid.
//
// If passed RequestVerifyData is nil, ErrNilRequestVerifyData returns.
func VerifyRequestData(src RequestVerifyData) error {
	if src == nil {
		return ErrNilRequestVerifyData
	}

	verSrc, err := GroupVerifyPayloads(
		src,
		src,
		NewVerifiedSessionToken(
			src.GetSessionToken(),
		),
		NewVerifiedBearerToken(
			src.GetBearerToken(),
		),
		ExtendedHeadersSignedData(src),
	)
	if err != nil {
		return err
	}

	return VerifyAccumulatedSignatures(verSrc)
}

// SignedData returns payload bytes concatenation from all sources keeping order.
func (s signSourceGroup) SignedData() ([]byte, error) {
	chunks := make([][]byte, 0, len(s.sources))
	sz := 0

	for i := range s.sources {
		data, err := s.sources[i].SignedData()
		if err != nil {
			return nil, errors.Wrapf(err, "could not get signed payload of element #%d", i)
		}

		chunks = append(chunks, data)

		sz += len(data)
	}

	res := make([]byte, sz)
	off := 0

	for i := range chunks {
		off += copy(res[off:], chunks[i])
	}

	return res, nil
}

// SignedData returns payload bytes concatenation from all readers.
func (s signReadersGroup) SignedData() ([]byte, error) {
	return SignedDataFromReader(s)
}

// SignedDataSize returns the sum of sizes of all readers.
func (s signReadersGroup) SignedDataSize() (sz int) {
	for i := range s.readers {
		sz += s.readers[i].SignedDataSize()
	}

	return
}

// ReadSignedData reads data from all readers to passed buffer keeping order.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (s signReadersGroup) ReadSignedData(p []byte) (int, error) {
	sz := s.SignedDataSize()
	if len(p) < sz {
		return 0, io.ErrUnexpectedEOF
	}

	off := 0

	for i := range s.readers {
		n, err := s.readers[i].ReadSignedData(p[off:])
		off += n
		if err != nil {
			return off, errors.Wrapf(err, "could not read signed payload of element #%d", i)
		}
	}

	return off, nil
}

// GroupSignedPayloads groups SignKeyPairAccumulator and SignedDataSource list to DataWithSignKeyAccumulator.
//
// If passed SignKeyPairAccumulator is nil, ErrNilSignKeyPairAccumulator returns.
//
// Signed payload of the result is a concatenation of payloads of list elements keeping order.
// Nil elements in list are ignored.
//
// If all elements implement SignedDataReader, result implements it too.
func GroupSignedPayloads(acc SignKeyPairAccumulator, sources ...SignedDataSource) (DataWithSignKeyAccumulator, error) {
	if acc == nil {
		return nil, ErrNilSignKeyPairAccumulator
	}

	return groupPayloads(acc, nil, sources...), nil
}

// GroupVerifyPayloads groups SignKeyPairSource and SignedDataSource list to DataWithSignKeySource.
//
// If passed SignKeyPairSource is nil, ErrNilSignatureKeySource returns.
//
// Signed payload of the result is a concatenation of payloads of list elements keeping order.
// Nil elements in list are ignored.
//
// If all elements implement SignedDataReader, result implements it too.
func GroupVerifyPayloads(src SignKeyPairSource, sources ...SignedDataSource) (DataWithSignKeySource, error) {
	if src == nil {
		return nil, ErrNilSignatureKeySource
	}

	return groupPayloads(nil, src, sources...), nil
}

func groupPayloads(acc SignKeyPairAccumulator, src SignKeyPairSource, sources ...SignedDataSource) interface {
	SignedDataSource
	SignKeyPairSource
	SignKeyPairAccumulator
} {
	var allReaders bool

	for i := range sources {
		if sources[i] == nil {
			continue
		} else if _, allReaders = sources[i].(SignedDataReader); !allReaders {
			break
		}
	}

	if !allReaders {
		res := &signSourceGroup{
			SignKeyPairSource:      src,
			SignKeyPairAccumulator: acc,

			sources: make([]SignedDataSource, 0, len(sources)),
		}

		for i := range sources {
			if sources[i] != nil {
				res.sources = append(res.sources, sources[i])
			}
		}

		return res
	}

	res := &signReadersGroup{
		SignKeyPairSource:      src,
		SignKeyPairAccumulator: acc,

		readers: make([]SignedDataReader, 0, len(sources)),
	}

	for i := range sources {
		if sources[i] != nil {
			res.readers = append(res.readers, sources[i].(SignedDataReader))
		}
	}

	return res
}
