package service

import (
	"bytes"
	"log"
	"math"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/nspcc-dev/neofs-api-go/refs"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func BenchmarkSignRequestHeader(b *testing.B) {
	key := test.DecodeKey(0)

	custom := testCustomField{1, 2, 3, 4, 5, 6, 7, 8}

	some := &TestRequest{
		IntField:    math.MaxInt32,
		StringField: "TestRequestStringField",
		BytesField:  make([]byte, 1<<22),
		CustomField: &custom,
		RequestMetaHeader: RequestMetaHeader{
			TTL:   math.MaxInt32 - 8,
			Epoch: math.MaxInt64 - 12,
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		require.NoError(b, SignRequestHeader(key, some))
	}
}

func BenchmarkVerifyRequestHeader(b *testing.B) {
	custom := testCustomField{1, 2, 3, 4, 5, 6, 7, 8}

	some := &TestRequest{
		IntField:    math.MaxInt32,
		StringField: "TestRequestStringField",
		BytesField:  make([]byte, 1<<22),
		CustomField: &custom,
		RequestMetaHeader: RequestMetaHeader{
			TTL:   math.MaxInt32 - 8,
			Epoch: math.MaxInt64 - 12,
		},
	}

	for i := 0; i < 10; i++ {
		key := test.DecodeKey(i)
		require.NoError(b, SignRequestHeader(key, some))
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		require.NoError(b, VerifyRequestHeader(some))
	}
}

func TestSignRequestHeader(t *testing.T) {
	req := &TestRequest{
		IntField:    math.MaxInt32,
		StringField: "TestRequestStringField",
		BytesField:  []byte("TestRequestBytesField"),
	}

	key := test.DecodeKey(0)
	peer := crypto.MarshalPublicKey(&key.PublicKey)

	data, err := req.Marshal()
	require.NoError(t, err)

	require.NoError(t, SignRequestHeader(key, req))

	require.Len(t, req.Signatures, 1)
	for i := range req.Signatures {
		sign := req.Signatures[i].GetSign()
		require.Equal(t, peer, req.Signatures[i].GetPeer())
		require.NoError(t, crypto.Verify(&key.PublicKey, data, sign))
	}
}

func TestVerifyRequestHeader(t *testing.T) {
	req := &TestRequest{
		IntField:          math.MaxInt32,
		StringField:       "TestRequestStringField",
		BytesField:        []byte("TestRequestBytesField"),
		RequestMetaHeader: RequestMetaHeader{TTL: 10},
	}

	for i := 0; i < 10; i++ {
		req.TTL--
		require.NoError(t, SignRequestHeader(test.DecodeKey(i), req))
	}

	require.NoError(t, VerifyRequestHeader(req))
}

func TestMaintainableRequest(t *testing.T) {
	req := &TestRequest{
		IntField:          math.MaxInt32,
		StringField:       "TestRequestStringField",
		BytesField:        []byte("TestRequestBytesField"),
		RequestMetaHeader: RequestMetaHeader{TTL: 10},
	}

	count := 10
	owner := test.DecodeKey(count + 1)

	for i := 0; i < count; i++ {
		req.TTL--

		key := test.DecodeKey(i)
		require.NoError(t, SignRequestHeader(key, req))

		// sign first key (session key) by owner key
		if i == 0 {
			sign, err := crypto.Sign(owner, crypto.MarshalPublicKey(&key.PublicKey))
			require.NoError(t, err)

			req.SetOwner(&owner.PublicKey, sign)
		}
	}

	{ // Validate owner
		user, err := refs.NewOwnerID(&owner.PublicKey)
		require.NoError(t, err)
		require.NoError(t, req.CheckOwner(user))
	}

	{ // Good case:
		require.NoError(t, VerifyRequestHeader(req))

		// validate, that first key (session key) was signed with owner
		signatures := req.GetSignatures()

		require.Len(t, signatures, count)

		pub, err := req.GetOwner()
		require.NoError(t, err)

		require.Equal(t, &owner.PublicKey, pub)
	}

	{ // wrong owner:
		req.Signatures[0].Origin = nil

		pub, err := req.GetOwner()
		require.NoError(t, err)

		require.NotEqual(t, &owner.PublicKey, pub)
	}

	{ // Wrong signatures:
		copy(req.Signatures[count-1].Sign, req.Signatures[count-1].Peer)
		err := VerifyRequestHeader(req)
		require.EqualError(t, errors.Cause(err), crypto.ErrInvalidSignature.Error())
	}
}

func TestVerifyAndSignRequestHeaderWithoutCloning(t *testing.T) {
	key := test.DecodeKey(0)

	custom := testCustomField{1, 2, 3, 4, 5, 6, 7, 8}

	b := &TestRequest{
		IntField:    math.MaxInt32,
		StringField: "TestRequestStringField",
		BytesField:  []byte("TestRequestBytesField"),
		CustomField: &custom,
		RequestMetaHeader: RequestMetaHeader{
			TTL:   math.MaxInt32 - 8,
			Epoch: math.MaxInt64 - 12,
		},
	}

	require.NoError(t, SignRequestHeader(key, b))
	require.NoError(t, VerifyRequestHeader(b))

	require.Len(t, b.Signatures, 1)
	require.Equal(t, custom, *b.CustomField)
	require.Equal(t, uint32(math.MaxInt32-8), b.GetTTL())
	require.Equal(t, uint64(math.MaxInt64-12), b.GetEpoch())

	buf := bytes.NewBuffer(nil)
	log.SetOutput(buf)

	cp, ok := proto.Clone(b).(*TestRequest)
	require.True(t, ok)
	require.NotEqual(t, b, cp)

	require.Contains(t, buf.String(), "proto: don't know how to copy")
}
