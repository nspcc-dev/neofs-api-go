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
	err   error
	data  []byte
	sig   []byte
	key   *ecdsa.PublicKey
	token SessionToken

	bearer BearerToken

	extHdrs []ExtendedHeader
}

type testSignedDataReader struct {
	*testSignedDataSrc
}

func (s testSignedDataSrc) GetSignature() []byte {
	return s.sig
}

func (s testSignedDataSrc) GetSignKeyPairs() []SignKeyPair {
	return []SignKeyPair{
		newSignatureKeyPair(s.key, s.sig),
	}
}

func (s testSignedDataSrc) SignedData() ([]byte, error) {
	return s.data, s.err
}

func (s *testSignedDataSrc) AddSignKey(sig []byte, key *ecdsa.PublicKey) {
	s.key = key
	s.sig = sig
}

func testData(t *testing.T, sz int) []byte {
	d := make([]byte, sz)
	_, err := rand.Read(d)
	require.NoError(t, err)
	return d
}

func (s testSignedDataSrc) GetSessionToken() SessionToken {
	return s.token
}

func (s testSignedDataSrc) GetBearerToken() BearerToken {
	return s.bearer
}

func (s testSignedDataSrc) ExtendedHeaders() []ExtendedHeader {
	return s.extHdrs
}

func (s testSignedDataReader) SignedDataSize() int {
	return len(s.data)
}

func (s testSignedDataReader) ReadSignedData(buf []byte) (int, error) {
	if s.err != nil {
		return 0, s.err
	}

	var err error
	if len(buf) < len(s.data) {
		err = io.ErrUnexpectedEOF
	}
	return copy(buf, s.data), err
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
			data: testData(t, 10),
		}

		// create custom error for data source
		src.err = errors.New("test error for data source")

		_, err = DataSignature(sk, src)
		require.EqualError(t, err, src.err.Error())

		// reset error to nil
		src.err = nil

		// calculate data signature
		sig, err := DataSignature(sk, src)
		require.NoError(t, err)

		// ascertain that the signature passes verification
		require.NoError(t, crypto.Verify(&sk.PublicKey, src.data, sig))
	})

	t.Run("signed data reader", func(t *testing.T) {
		// create test signed data reader
		src := &testSignedDataSrc{
			data: testData(t, 10),
		}

		// create custom error for signed data reader
		src.err = errors.New("test error for signed data reader")

		sig, err := DataSignature(sk, src)
		require.EqualError(t, err, src.err.Error())

		// reset error to nil
		src.err = nil

		// calculate data signature
		sig, err = DataSignature(sk, src)
		require.NoError(t, err)

		// ascertain that the signature passes verification
		require.NoError(t, crypto.Verify(&sk.PublicKey, src.data, sig))
	})
}

func TestAddSignatureWithKey(t *testing.T) {
	require.NoError(t,
		AddSignatureWithKey(
			test.DecodeKey(0),
			&testSignedDataSrc{
				data: testData(t, 10),
			},
		),
	)
}

func TestVerifySignatures(t *testing.T) {
	// empty signatures
	require.NoError(t, VerifySignatures(nil))

	// create test signature source
	src := &testSignedDataSrc{
		data: testData(t, 10),
	}

	// create private key for test
	sk := test.DecodeKey(0)

	// calculate a signature of the data
	sig, err := crypto.Sign(sk, src.data)
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
	src := &testSignedDataSrc{
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
	src := &testSignedDataSrc{
		data: testData(t, 10),
	}

	// nil public key
	require.EqualError(t,
		VerifySignatureWithKey(nil, src),
		crypto.ErrEmptyPublicKey.Error(),
	)

	// create test private key
	sk := test.DecodeKey(0)

	var err error

	// calculate a signature
	src.sig, err = crypto.Sign(sk, src.data)
	require.NoError(t, err)

	// ascertain that verification is passed
	require.NoError(t, VerifySignatureWithKey(&sk.PublicKey, src))

	// break the signature
	src.sig[0]++

	// ascertain that verification is failed
	require.Error(t, VerifySignatureWithKey(&sk.PublicKey, src))
}

func TestSignVerifyRequestData(t *testing.T) {
	// sign with empty RequestSignedData
	require.EqualError(t,
		SignRequestData(nil, nil),
		ErrNilRequestSignedData.Error(),
	)

	// verify with empty RequestVerifyData
	require.EqualError(t,
		VerifyRequestData(nil),
		ErrNilRequestVerifyData.Error(),
	)

	// create test session token
	var (
		token    = new(Token)
		initVerb = Token_Info_Verb(1)

		bearer      = wrapBearerTokenMsg(new(BearerTokenMsg))
		bearerEpoch = uint64(8)

		extHdrKey = "key"
		extHdr    = new(RequestExtendedHeader_KV)
	)

	token.SetVerb(initVerb)

	bearer.SetExpirationEpoch(bearerEpoch)

	extHdr.SetK(extHdrKey)

	// create test data with token
	src := &testSignedDataSrc{
		data:  testData(t, 10),
		token: token,

		bearer: bearer,

		extHdrs: []ExtendedHeader{
			wrapExtendedHeaderKV(extHdr),
		},
	}

	// create test private key
	sk := test.DecodeKey(0)

	// sign with private key
	require.NoError(t, SignRequestData(sk, src))

	// ascertain that verification is passed
	require.NoError(t, VerifyRequestData(src))

	// break the data
	src.data[0]++

	// ascertain that verification is failed
	require.Error(t, VerifyRequestData(src))

	// restore the data
	src.data[0]--

	// break the token
	token.SetVerb(initVerb + 1)

	// ascertain that verification is failed
	require.Error(t, VerifyRequestData(src))

	// restore the token
	token.SetVerb(initVerb)

	// ascertain that verification is passed
	require.NoError(t, VerifyRequestData(src))

	// break the Bearer token
	bearer.SetExpirationEpoch(bearerEpoch + 1)

	// ascertain that verification is failed
	require.Error(t, VerifyRequestData(src))

	// restore the Bearer token
	bearer.SetExpirationEpoch(bearerEpoch)

	// ascertain that verification is passed
	require.NoError(t, VerifyRequestData(src))

	// break the extended header
	extHdr.SetK(extHdrKey + "1")

	// ascertain that verification is failed
	require.Error(t, VerifyRequestData(src))

	// restore the extended header
	extHdr.SetK(extHdrKey)

	// ascertain that verification is passed
	require.NoError(t, VerifyRequestData(src))

	// wrap to data reader
	rdr := &testSignedDataReader{
		testSignedDataSrc: src,
	}

	// sign with private key
	require.NoError(t, SignRequestData(sk, rdr))

	// ascertain that verification is passed
	require.NoError(t, VerifyRequestData(rdr))
}
