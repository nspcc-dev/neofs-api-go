package object_test

import (
	"fmt"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/object"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/object/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
	"github.com/stretchr/testify/require"
)

func TestShortHeader_StableMarshal(t *testing.T) {
	hdrFrom := generateShortHeader("Owner ID")
	transport := new(grpc.ShortHeader)

	t.Run("non empty", func(t *testing.T) {
		wire, err := hdrFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		hdrTo := object.ShortHeaderFromGRPCMessage(transport)
		require.Equal(t, hdrFrom, hdrTo)
	})
}

func TestAttribute_StableMarshal(t *testing.T) {
	from := generateAttribute("Key", "Value")
	transport := new(grpc.Header_Attribute)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.AttributeFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestSplitHeader_StableMarshal(t *testing.T) {
	from := generateSplit("Split Outside")
	hdr := generateHeader(123)
	from.SetParentHeader(hdr)

	transport := new(grpc.Header_Split)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.SplitHeaderFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestHeader_StableMarshal(t *testing.T) {
	insideHeader := generateHeader(100)
	split := generateSplit("Split")
	split.SetParentHeader(insideHeader)

	from := generateHeader(500)
	from.SetSplit(split)

	transport := new(grpc.Header)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.HeaderFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestObject_StableMarshal(t *testing.T) {
	from := generateObject("Payload")
	transport := new(grpc.Object)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.ObjectFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestGetRequestBody_StableMarshal(t *testing.T) {
	from := generateGetRequestBody("Container ID", "Object ID")
	transport := new(grpc.GetRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.GetRequestBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestGetResponseBody_StableMarshal(t *testing.T) {
	initFrom := generateGetResponseBody(true)
	chunkFrom := generateGetResponseBody(false)
	transport := new(grpc.GetResponse_Body)

	t.Run("init non empty", func(t *testing.T) {
		wire, err := initFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.GetResponseBodyFromGRPCMessage(transport)
		require.Equal(t, initFrom, to)
	})

	t.Run("chunk non empty", func(t *testing.T) {
		wire, err := chunkFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.GetResponseBodyFromGRPCMessage(transport)
		require.Equal(t, chunkFrom, to)
	})
}

func TestPutRequestBody_StableMarshal(t *testing.T) {
	initFrom := generatePutRequestBody(true)
	chunkFrom := generatePutRequestBody(false)
	transport := new(grpc.PutRequest_Body)

	t.Run("init non empty", func(t *testing.T) {
		wire, err := initFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.PutRequestBodyFromGRPCMessage(transport)
		require.Equal(t, initFrom, to)
	})

	t.Run("chunk non empty", func(t *testing.T) {
		wire, err := chunkFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.PutRequestBodyFromGRPCMessage(transport)
		require.Equal(t, chunkFrom, to)
	})
}

func TestPutRequestBody_StableSize(t *testing.T) {
	from := generatePutResponseBody("Object ID")
	transport := new(grpc.PutResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.PutResponseBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestDeleteRequestBody_StableMarshal(t *testing.T) {
	from := generateDeleteRequestBody("Container ID", "Object ID")
	transport := new(grpc.DeleteRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.DeleteRequestBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestDeleteResponseBody_StableMarshal(t *testing.T) {
	from := generateDeleteResponseBody()
	transport := new(grpc.DeleteResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.DeleteResponseBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestSplitHeaderFromGRPCMessage(t *testing.T) {
	from := generateHeadRequestBody("Container ID", "Object ID")
	transport := new(grpc.HeadRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.HeadRequestBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestHeadResponseBody_StableMarshal(t *testing.T) {
	shortFrom := generateHeadResponseBody(true)
	fullFrom := generateHeadResponseBody(false)
	transport := new(grpc.HeadResponse_Body)

	t.Run("short header non empty", func(t *testing.T) {
		wire, err := shortFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.HeadResponseBodyFromGRPCMessage(transport)
		require.Equal(t, shortFrom, to)
	})

	t.Run("full header non empty", func(t *testing.T) {
		wire, err := fullFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.HeadResponseBodyFromGRPCMessage(transport)
		require.Equal(t, fullFrom, to)
	})
}

func TestSearchRequestBody_StableMarshal(t *testing.T) {
	from := generateSearchRequestBody(10, "Container ID")
	transport := new(grpc.SearchRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.SearchRequestBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestSearchResponseBody_StableMarshal(t *testing.T) {
	from := generateSearchResponseBody(10)
	transport := new(grpc.SearchResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.SearchResponseBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestGetRangeRequestBody_StableMarshal(t *testing.T) {
	from := generateRangeRequestBody("Container ID", "Object ID")
	transport := new(grpc.GetRangeRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.GetRangeRequestBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestGetRangeResponseBody_StableMarshal(t *testing.T) {
	from := generateRangeResponseBody("some data")
	transport := new(grpc.GetRangeResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.GetRangeResponseBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestGetRangeHashRequestBody_StableMarshal(t *testing.T) {
	from := generateRangeHashRequestBody("Container ID", "Object ID", 5)
	transport := new(grpc.GetRangeHashRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.GetRangeHashRequestBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestGetRangeHashResponseBody_StableMarshal(t *testing.T) {
	from := generateRangeHashResponseBody(5)
	transport := new(grpc.GetRangeHashResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		to := object.GetRangeHashResponseBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func generateOwner(id string) *refs.OwnerID {
	owner := new(refs.OwnerID)
	owner.SetValue([]byte(id))

	return owner
}

func generateObjectID(id string) *refs.ObjectID {
	oid := new(refs.ObjectID)
	oid.SetValue([]byte(id))

	return oid
}

func generateContainerID(id string) *refs.ContainerID {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte(id))

	return cid
}

func generateSignature(k, v string) *refs.Signature {
	sig := new(refs.Signature)
	sig.SetKey([]byte(k))
	sig.SetSign([]byte(v))

	return sig
}

func generateVersion(maj, min uint32) *refs.Version {
	version := new(refs.Version)
	version.SetMajor(maj)
	version.SetMinor(min)

	return version
}

func generateAddress(cid, oid string) *refs.Address {
	addr := new(refs.Address)
	addr.SetObjectID(generateObjectID(oid))
	addr.SetContainerID(generateContainerID(cid))

	return addr
}

func generateSessionToken(id string) *session.SessionToken {
	lifetime := new(session.TokenLifetime)
	lifetime.SetExp(1)
	lifetime.SetNbf(2)
	lifetime.SetIat(3)

	objectCtx := new(session.ObjectSessionContext)
	objectCtx.SetVerb(session.ObjectVerbPut)
	objectCtx.SetAddress(generateAddress("Container ID", "Object ID"))

	tokenBody := new(session.SessionTokenBody)
	tokenBody.SetID([]byte(id))
	tokenBody.SetOwnerID(generateOwner("Owner ID"))
	tokenBody.SetSessionKey([]byte(id))
	tokenBody.SetLifetime(lifetime)
	tokenBody.SetContext(objectCtx)

	sessionToken := new(session.SessionToken)
	sessionToken.SetBody(tokenBody)
	sessionToken.SetSignature(generateSignature("public key", id))

	return sessionToken
}

func generateShortHeader(id string) *object.ShortHeader {
	hdr := new(object.ShortHeader)
	hdr.SetOwnerID(generateOwner(id))
	hdr.SetVersion(generateVersion(2, 0))
	hdr.SetCreationEpoch(200)
	hdr.SetObjectType(object.TypeRegular)
	hdr.SetPayloadLength(10)

	return hdr
}

func generateAttribute(k, v string) *object.Attribute {
	attr := new(object.Attribute)
	attr.SetValue(v)
	attr.SetKey(k)

	return attr
}

func generateSplit(sig string) *object.SplitHeader {
	split := new(object.SplitHeader)
	split.SetChildren([]*refs.ObjectID{
		generateObjectID("Child 1"),
		generateObjectID("Child 2"),
	})
	split.SetParent(generateObjectID("Parent"))
	split.SetParentSignature(generateSignature("Key", sig))
	split.SetPrevious(generateObjectID("Previous"))

	return split
}

func generateChecksum(data string) *refs.Checksum {
	checksum := new(refs.Checksum)
	checksum.SetType(refs.TillichZemor)
	checksum.SetSum([]byte(data))

	return checksum
}

func generateHeader(ln uint64) *object.Header {
	hdr := new(object.Header)
	hdr.SetPayloadLength(ln)
	hdr.SetCreationEpoch(ln / 2)
	hdr.SetVersion(generateVersion(2, 0))
	hdr.SetOwnerID(generateOwner("Owner ID"))
	hdr.SetContainerID(generateContainerID("Contanier ID"))
	hdr.SetAttributes([]*object.Attribute{
		generateAttribute("One", "Two"),
		generateAttribute("Three", "Four"),
	})
	hdr.SetHomomorphicHash(generateChecksum("Homomorphic Hash"))
	hdr.SetObjectType(object.TypeRegular)
	hdr.SetPayloadHash(generateChecksum("Payload Hash"))
	hdr.SetSessionToken(generateSessionToken(string(ln)))

	return hdr
}

func generateObject(data string) *object.Object {
	insideHeader := generateHeader(100)
	split := generateSplit("Split")
	split.SetParentHeader(insideHeader)

	outsideHeader := generateHeader(500)
	outsideHeader.SetSplit(split)

	obj := new(object.Object)
	obj.SetSignature(generateSignature("Public Key", "Signature"))
	obj.SetObjectID(generateObjectID("Object ID"))
	obj.SetPayload([]byte(data))
	obj.SetHeader(outsideHeader)

	return obj
}

func generateGetRequestBody(cid, oid string) *object.GetRequestBody {
	req := new(object.GetRequestBody)
	req.SetAddress(generateAddress(cid, oid))
	req.SetRaw(true)

	return req
}

func generateGetResponseBody(flag bool) *object.GetResponseBody {
	resp := new(object.GetResponseBody)
	var part object.GetObjectPart

	if flag {
		init := new(object.GetObjectPartInit)
		init.SetObjectID(generateObjectID("Object ID"))
		init.SetSignature(generateSignature("Key", "Signature"))
		init.SetHeader(generateHeader(10))
		part = init
	} else {
		chunk := new(object.GetObjectPartChunk)
		chunk.SetChunk([]byte("Some data chunk"))
		part = chunk
	}
	resp.SetObjectPart(part)

	return resp
}

func generatePutRequestBody(flag bool) *object.PutRequestBody {
	req := new(object.PutRequestBody)
	var part object.PutObjectPart

	if flag {
		init := new(object.PutObjectPartInit)
		init.SetObjectID(generateObjectID("Object ID"))
		init.SetSignature(generateSignature("Key", "Signature"))
		init.SetHeader(generateHeader(10))
		init.SetCopiesNumber(1)
		part = init
	} else {
		chunk := new(object.PutObjectPartChunk)
		chunk.SetChunk([]byte("Some data chunk"))
		part = chunk
	}
	req.SetObjectPart(part)

	return req
}

func generatePutResponseBody(oid string) *object.PutResponseBody {
	resp := new(object.PutResponseBody)
	resp.SetObjectID(generateObjectID(oid))

	return resp
}

func generateDeleteRequestBody(cid, oid string) *object.DeleteRequestBody {
	req := new(object.DeleteRequestBody)
	req.SetAddress(generateAddress(cid, oid))

	return req
}

func generateDeleteResponseBody() *object.DeleteResponseBody {
	return new(object.DeleteResponseBody)
}

func generateHeadRequestBody(cid, oid string) *object.HeadRequestBody {
	req := new(object.HeadRequestBody)
	req.SetAddress(generateAddress(cid, oid))
	req.SetRaw(true)
	req.SetMainOnly(true)

	return req
}

func generateHeadResponseBody(flag bool) *object.HeadResponseBody {
	req := new(object.HeadResponseBody)
	var part object.GetHeaderPart

	if flag {
		short := new(object.GetHeaderPartShort)
		short.SetShortHeader(generateShortHeader("short id"))
		part = short
	} else {
		full := new(object.GetHeaderPartFull)
		full.SetHeader(generateHeader(30))
		part = full
	}

	req.SetHeaderPart(part)

	return req
}

func generateFilter(k, v string) *object.SearchFilter {
	f := new(object.SearchFilter)
	f.SetName(k)
	f.SetValue(v)
	f.SetMatchType(object.MatchStringEqual)

	return f
}

func generateSearchRequestBody(n int, id string) *object.SearchRequestBody {
	req := new(object.SearchRequestBody)
	req.SetContainerID(generateContainerID(id))
	req.SetVersion(1)

	ff := make([]*object.SearchFilter, n)

	for i := 0; i < n; i++ {
		ff[i] = generateFilter("Some Key", fmt.Sprintf("Value %d", i+1))
	}
	req.SetFilters(ff)

	return req
}

func generateSearchResponseBody(n int) *object.SearchResponseBody {
	resp := new(object.SearchResponseBody)
	list := make([]*refs.ObjectID, n)
	for i := 0; i < n; i++ {
		list[i] = generateObjectID(fmt.Sprintf("Object ID %d", i+1))
	}

	resp.SetIDList(list)

	return resp
}

func generateRange(off, ln uint64) *object.Range {
	r := new(object.Range)
	r.SetOffset(off)
	r.SetLength(ln)

	return r
}

func generateRangeRequestBody(cid, oid string) *object.GetRangeRequestBody {
	req := new(object.GetRangeRequestBody)
	req.SetAddress(generateAddress(cid, oid))
	req.SetRange(generateRange(10, 20))

	return req
}

func generateRangeResponseBody(data string) *object.GetRangeResponseBody {
	resp := new(object.GetRangeResponseBody)
	resp.SetChunk([]byte(data))

	return resp
}

func generateRangeHashRequestBody(cid, oid string, n int) *object.GetRangeHashRequestBody {
	req := new(object.GetRangeHashRequestBody)
	req.SetAddress(generateAddress(cid, oid))

	rngs := make([]*object.Range, n)
	for i := 0; i < n; i++ {
		rngs[i] = generateRange(100, 200+uint64(n))
	}

	req.SetRanges(rngs)
	req.SetSalt([]byte("xor salt"))
	req.SetType(refs.TillichZemor)

	return req
}

func generateRangeHashResponseBody(n int) *object.GetRangeHashResponseBody {
	resp := new(object.GetRangeHashResponseBody)

	list := make([][]byte, n)
	for i := 0; i < n; i++ {
		list[i] = []byte("Some homomorphic hash data" + string(n))
	}

	resp.SetType(refs.TillichZemor)
	resp.SetHashList(list)

	return resp
}
