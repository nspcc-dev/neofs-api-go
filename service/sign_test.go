package service

import (
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

	// nil data source
	_, err = DataSignature(nil, nil)
	require.EqualError(t, err, ErrNilSignedDataSource.Error())

	// nil private key
	_, err = DataSignature(new(testSignedDataSrc), nil)
	require.EqualError(t, err, crypto.ErrEmptyPrivateKey.Error())

	// create test private key
	sk := test.DecodeKey(0)

	t.Run("common signed data source", func(t *testing.T) {
		// create test data source
		src := &testSignedDataSrc{
			d: testData(t, 10),
		}

		// create custom error for data source
		src.e = errors.New("test error for data source")

		_, err = DataSignature(src, sk)
		require.EqualError(t, err, src.e.Error())

		// reset error to nil
		src.e = nil

		// calculate data signature
		sig, err := DataSignature(src, sk)
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

		sig, err := DataSignature(src, sk)
		require.EqualError(t, err, src.e.Error())

		// reset error to nil
		src.e = nil

		// calculate data signature
		sig, err = DataSignature(src, sk)
		require.NoError(t, err)

		// ascertain that the signature passes verification
		require.NoError(t, crypto.Verify(&sk.PublicKey, src.d, sig))
	})
}
