package service

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"io"
	"testing"

	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

type testSignedDataSrc struct {
	e error
	d []byte
}

type testSignedDataReader struct {
	SignedDataSource

	e error
	d []byte
}

type testKeySigAccum struct {
	data []byte
	sig  []byte
	key  *ecdsa.PublicKey
}

func (s testKeySigAccum) GetSignature() []byte {
	return s.sig
}

func (s testKeySigAccum) GetSignKeyPairs() []SignKeyPair {
	return []SignKeyPair{
		newSignatureKeyPair(s.key, s.sig),
	}
}

func (s testKeySigAccum) SignedData() ([]byte, error) {
	return s.data, nil
}

func (s testKeySigAccum) AddSignKey(sig []byte, key *ecdsa.PublicKey) {
	s.key = key
	s.sig = sig
}

func testData(t *testing.T, sz int) []byte {
	d := make([]byte, sz)
	_, err := rand.Read(d)
	require.NoError(t, err)
	return d
}

func (s testSignedDataReader) SignedDataSize() int {
	return len(s.d)
}

func (s testSignedDataReader) ReadSignedData(buf []byte) (int, error) {
	if s.e != nil {
		return 0, s.e
	}

	var err error
	if len(buf) < len(s.d) {
		err = io.ErrUnexpectedEOF
	}
	return copy(buf, s.d), err
}

func (s testSignedDataSrc) SignedData() ([]byte, error) {
	return s.d, s.e
}

func TestDataSignature(t *testing.T) {
	var err error

	// nil private key
	_, err = DataSignature(nil, nil)
	require.EqualError(t, err, crypto.ErrEmptyPrivateKey.Error())

	// create test private key
	sk := test.DecodeKey(0)

	// nil private key
	_, err = DataSignature(sk, nil)
	require.EqualError(t, err, ErrNilSignedDataSource.Error())

	t.Run("common signed data source", func(t *testing.T) {
		// create test data source
		src := &testSignedDataSrc{
			d: testData(t, 10),
		}

		// create custom error for data source
		src.e = errors.New("test error for data source")

		_, err = DataSignature(sk, src)
		require.EqualError(t, err, src.e.Error())

		// reset error to nil
		src.e = nil

		// calculate data signature
		sig, err := DataSignature(sk, src)
		require.NoError(t, err)

		// ascertain that the signature passes verification
		require.NoError(t, crypto.Verify(&sk.PublicKey, src.d, sig))
	})

	t.Run("signed data reader", func(t *testing.T) {
		// create test signed data reader
		src := &testSignedDataReader{
			d: testData(t, 10),
		}

		// create custom error for signed data reader
		src.e = errors.New("test error for signed data reader")

		sig, err := DataSignature(sk, src)
		require.EqualError(t, err, src.e.Error())

		// reset error to nil
		src.e = nil

		// calculate data signature
		sig, err = DataSignature(sk, src)
		require.NoError(t, err)

		// ascertain that the signature passes verification
		require.NoError(t, crypto.Verify(&sk.PublicKey, src.d, sig))
	})
}

func TestAddSignatureWithKey(t *testing.T) {
	// create test data
	data := testData(t, 10)

	// create test private key
	sk := test.DecodeKey(0)

	// create test signature accumulator
	var s SignatureKeyAccumulator = &testKeySigAccum{
		data: data,
	}

	require.NoError(t, AddSignatureWithKey(s, sk))
}

func TestVerifySignatures(t *testing.T) {
	// empty signatures
	require.NoError(t, VerifySignatures(nil))

	// create test signature source
	src := &testSignedDataSrc{
		d: testData(t, 10),
	}

	// create private key for test
	sk := test.DecodeKey(0)

	// calculate a signature of the data
	sig, err := crypto.Sign(sk, src.d)
	require.NoError(t, err)

	// ascertain that verification is passed
	require.NoError(t,
		VerifySignatures(
			src,
			newSignatureKeyPair(&sk.PublicKey, sig),
		),
	)

	// break the signature
	sig[0]++

	require.Error(t,
		VerifySignatures(
			src,
			newSignatureKeyPair(&sk.PublicKey, sig),
		),
	)

	// restore the signature
	sig[0]--

	// empty data source
	require.EqualError(t,
		VerifySignatures(nil, nil),
		ErrNilSignedDataSource.Error(),
	)

}

func TestVerifyAccumulatedSignatures(t *testing.T) {
	// nil signature source
	require.EqualError(t,
		VerifyAccumulatedSignatures(nil),
		ErrNilSignatureKeySource.Error(),
	)

	// create test private key
	sk := test.DecodeKey(0)

	// create signature source
	src := &testKeySigAccum{
		data: testData(t, 10),
		key:  &sk.PublicKey,
	}

	var err error

	// calculate a signature
	src.sig, err = crypto.Sign(sk, src.data)
	require.NoError(t, err)

	// ascertain that verification is passed
	require.NoError(t, VerifyAccumulatedSignatures(src))

	// break the signature
	src.sig[0]++

	// ascertain that verification is failed
	require.Error(t, VerifyAccumulatedSignatures(src))
}

func TestVerifySignatureWithKey(t *testing.T) {
	// nil signature source
	require.EqualError(t,
		VerifySignatureWithKey(nil, nil),
		ErrEmptyDataWithSignature.Error(),
	)

	// create test signature source
	src := &testKeySigAccum{
		data: testData(t, 10),
	}

	// nil public key
	require.EqualError(t,
		VerifySignatureWithKey(src, nil),
		crypto.ErrEmptyPublicKey.Error(),
	)

	// create test private key
	sk := test.DecodeKey(0)

	var err error

	// calculate a signature
	src.sig, err = crypto.Sign(sk, src.data)
	require.NoError(t, err)

	// ascertain that verification is passed
	require.NoError(t, VerifySignatureWithKey(src, &sk.PublicKey))

	// break the signature
	src.sig[0]++

	// ascertain that verification is failed
	require.Error(t, VerifySignatureWithKey(src, &sk.PublicKey))
}
