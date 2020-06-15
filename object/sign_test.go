package object

import (
	"crypto/rand"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/service"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestSignVerifyRequests(t *testing.T) {
	sk := test.DecodeKey(0)

	type sigType interface {
		service.RequestData
		service.SignKeyPairAccumulator
		service.SignKeyPairSource
		SetToken(*Token)
		SetBearer(*service.BearerTokenMsg)
	}

	items := []struct {
		constructor    func() sigType
		payloadCorrupt []func(sigType)
	}{
		{ // PutRequest.PutHeader
			constructor: func() sigType {
				return MakePutRequestHeader(new(Object))
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					obj := s.(*PutRequest).GetR().(*PutRequest_Header).Header.GetObject()
					obj.SystemHeader.PayloadLength++
				},
			},
		},
		{ // PutRequest.Chunk
			constructor: func() sigType {
				return MakePutRequestChunk(make([]byte, 10))
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					h := s.(*PutRequest).GetR().(*PutRequest_Chunk)
					h.Chunk[0]++
				},
			},
		},
		{ // GetRequest
			constructor: func() sigType {
				return new(GetRequest)
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					s.(*GetRequest).Address.CID[0]++
				},
				func(s sigType) {
					s.(*GetRequest).Address.ObjectID[0]++
				},
			},
		},
		{ // HeadRequest
			constructor: func() sigType {
				return new(HeadRequest)
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					s.(*HeadRequest).Address.CID[0]++
				},
				func(s sigType) {
					s.(*HeadRequest).Address.ObjectID[0]++
				},
				func(s sigType) {
					s.(*HeadRequest).FullHeaders = true
				},
			},
		},
		{ // DeleteRequest
			constructor: func() sigType {
				return new(DeleteRequest)
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					s.(*DeleteRequest).OwnerID[0]++
				},
				func(s sigType) {
					s.(*DeleteRequest).Address.CID[0]++
				},
				func(s sigType) {
					s.(*DeleteRequest).Address.ObjectID[0]++
				},
			},
		},
		{ // GetRangeRequest
			constructor: func() sigType {
				return new(GetRangeRequest)
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					s.(*GetRangeRequest).Range.Length++
				},
				func(s sigType) {
					s.(*GetRangeRequest).Range.Offset++
				},
				func(s sigType) {
					s.(*GetRangeRequest).Address.CID[0]++
				},
				func(s sigType) {
					s.(*GetRangeRequest).Address.ObjectID[0]++
				},
			},
		},
		{ // GetRangeHashRequest
			constructor: func() sigType {
				return &GetRangeHashRequest{
					Ranges: []Range{{}},
					Salt:   []byte{1, 2, 3},
				}
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					s.(*GetRangeHashRequest).Address.CID[0]++
				},
				func(s sigType) {
					s.(*GetRangeHashRequest).Address.ObjectID[0]++
				},
				func(s sigType) {
					s.(*GetRangeHashRequest).Salt[0]++
				},
				func(s sigType) {
					s.(*GetRangeHashRequest).Ranges[0].Length++
				},
				func(s sigType) {
					s.(*GetRangeHashRequest).Ranges[0].Offset++
				},
				func(s sigType) {
					s.(*GetRangeHashRequest).Ranges = nil
				},
			},
		},
		{ // GetRangeHashRequest
			constructor: func() sigType {
				return &SearchRequest{
					Query: []byte{1, 2, 3},
				}
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					s.(*SearchRequest).ContainerID[0]++
				},
				func(s sigType) {
					s.(*SearchRequest).Query[0]++
				},
				func(s sigType) {
					s.(*SearchRequest).QueryVersion++
				},
			},
		},
	}

	for _, item := range items {
		{ // token corruptions
			v := item.constructor()

			token := new(Token)
			v.SetToken(token)

			require.NoError(t, service.SignRequestData(sk, v))

			require.NoError(t, service.VerifyRequestData(v))

			token.SetSessionKey(append(token.GetSessionKey(), 1))

			require.Error(t, service.VerifyRequestData(v))
		}

		{ // Bearer token corruptions
			v := item.constructor()

			token := new(service.BearerTokenMsg)
			v.SetBearer(token)

			require.NoError(t, service.SignRequestData(sk, v))

			require.NoError(t, service.VerifyRequestData(v))

			token.SetACLRules(append(token.GetACLRules(), 1))

			require.Error(t, service.VerifyRequestData(v))
		}

		{ // payload corruptions
			for _, corruption := range item.payloadCorrupt {
				v := item.constructor()

				require.NoError(t, service.SignRequestData(sk, v))

				require.NoError(t, service.VerifyRequestData(v))

				corruption(v)

				require.Error(t, service.VerifyRequestData(v))
			}
		}
	}
}

func TestHeadRequest_ReadSignedData(t *testing.T) {
	t.Run("full headers", func(t *testing.T) {
		req := new(HeadRequest)

		// unset FullHeaders flag
		req.SetFullHeaders(false)

		// allocate two different buffers for reading
		buf1 := testData(t, req.SignedDataSize())
		buf2 := testData(t, req.SignedDataSize())

		// read to both buffers
		n1, err := req.ReadSignedData(buf1)
		require.NoError(t, err)

		n2, err := req.ReadSignedData(buf2)
		require.NoError(t, err)

		require.Equal(t, buf1[:n1], buf2[:n2])
	})
}

func testData(t *testing.T, sz int) []byte {
	data := make([]byte, sz)

	_, err := rand.Read(data)
	require.NoError(t, err)

	return data
}

func TestIntegrityHeaderSignMethods(t *testing.T) {
	// create new IntegrityHeader
	s := new(IntegrityHeader)

	// set test headers checksum
	s.SetHeadersChecksum([]byte{1, 2, 3})

	data, err := s.SignedData()
	require.NoError(t, err)
	require.Equal(t, data, s.GetHeadersChecksum())

	// add signature
	sig := []byte{4, 5, 6}
	s.AddSignKey(sig, nil)

	require.Equal(t, sig, s.GetSignature())
}
