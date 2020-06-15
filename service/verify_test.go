package service

import (
	"encoding/binary"
	"io"
	"math"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func (m TestRequest) SignedData() ([]byte, error) {
	return SignedDataFromReader(m)
}

func (m TestRequest) SignedDataSize() (sz int) {
	sz += 4

	sz += len(m.StringField)

	sz += len(m.BytesField)

	sz += m.CustomField.Size()

	return
}

func (m TestRequest) ReadSignedData(p []byte) (int, error) {
	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	binary.BigEndian.PutUint32(p[off:], uint32(m.IntField))
	off += 4

	off += copy(p[off:], []byte(m.StringField))

	off += copy(p[off:], m.BytesField)

	n, err := m.CustomField.MarshalTo(p[off:])
	off += n

	return off, err
}

func BenchmarkSignDataWithSessionToken(b *testing.B) {
	key := test.DecodeKey(0)

	customField := testCustomField{1, 2, 3, 4, 5, 6, 7, 8}

	token := new(Token)

	req := &TestRequest{
		IntField:    math.MaxInt32,
		StringField: "TestRequestStringField",
		BytesField:  make([]byte, 1<<22),
		CustomField: &customField,
	}

	req.SetTTL(math.MaxInt32 - 8)
	req.SetEpoch(math.MaxInt64 - 12)
	req.SetToken(token)
	req.SetBearer(new(BearerTokenMsg))

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		require.NoError(b, SignRequestData(key, req))
	}
}

func BenchmarkVerifyAccumulatedSignaturesWithToken(b *testing.B) {
	customField := testCustomField{1, 2, 3, 4, 5, 6, 7, 8}

	token := new(Token)

	req := &TestRequest{
		IntField:    math.MaxInt32,
		StringField: "TestRequestStringField",
		BytesField:  make([]byte, 1<<22),
		CustomField: &customField,
	}

	req.SetTTL(math.MaxInt32 - 8)
	req.SetEpoch(math.MaxInt64 - 12)
	req.SetToken(token)
	req.SetBearer(new(BearerTokenMsg))

	for i := 0; i < 10; i++ {
		key := test.DecodeKey(i)
		require.NoError(b, SignRequestData(key, req))
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		require.NoError(b, VerifyRequestData(req))
	}
}

func TestRequestVerificationHeader_SetToken(t *testing.T) {
	id, err := refs.NewUUID()
	require.NoError(t, err)

	token := new(Token)
	token.SetID(id)

	h := new(RequestVerificationHeader)

	h.SetToken(token)

	require.Equal(t, token, h.GetToken())
}

func TestRequestVerificationHeader_SetBearer(t *testing.T) {
	aclRules := []byte{1, 2, 3}

	token := new(BearerTokenMsg)
	token.SetACLRules(aclRules)

	h := new(RequestVerificationHeader)

	h.SetBearer(token)

	require.Equal(t, token, h.GetBearer())
}
