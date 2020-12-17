package object_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/object"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/object/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
	"github.com/stretchr/testify/require"
	goproto "google.golang.org/protobuf/proto"
)

func TestShortHeader_StableMarshal(t *testing.T) {
	hdrFrom := generateShortHeader("Owner ID")

	t.Run("non empty", func(t *testing.T) {
		wire, err := hdrFrom.StableMarshal(nil)
		require.NoError(t, err)

		hdrTo := new(object.ShortHeader)
		require.NoError(t, hdrTo.Unmarshal(wire))

		require.Equal(t, hdrFrom, hdrTo)
	})
}

func TestAttribute_StableMarshal(t *testing.T) {
	from := generateAttribute("Key", "Value")

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		to := new(object.Attribute)
		require.NoError(t, to.Unmarshal(wire))

		require.Equal(t, from, to)
	})
}

func TestSplitHeader_StableMarshal(t *testing.T) {
	from := generateSplit("Split Outside")
	hdr := generateHeader(123)
	from.SetParentHeader(hdr)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		to := new(object.SplitHeader)
		require.NoError(t, to.Unmarshal(wire))

		require.Equal(t, from, to)
	})
}

func TestHeader_StableMarshal(t *testing.T) {
	insideHeader := generateHeader(100)
	split := generateSplit("Split")
	split.SetParentHeader(insideHeader)

	from := generateHeader(500)
	from.SetSplit(split)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		to := new(object.Header)
		require.NoError(t, to.Unmarshal(wire))

		require.Equal(t, from, to)
	})
}

func TestObject_StableMarshal(t *testing.T) {
	from := generateObject("Payload")

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		to := new(object.Object)
		require.NoError(t, to.Unmarshal(wire))

		require.Equal(t, from, to)
	})
}

func TestGetRequestBody_StableMarshal(t *testing.T) {
	from := generateGetRequestBody("Container ID", "Object ID")
	transport := new(grpc.GetRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.GetRequestBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestGetResponseBody_StableMarshal(t *testing.T) {
	initFrom := generateGetResponseBody(0)
	chunkFrom := generateGetResponseBody(1)
	splitInfoFrom := generateGetResponseBody(2)
	transport := new(grpc.GetResponse_Body)

	t.Run("init non empty", func(t *testing.T) {
		wire, err := initFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.GetResponseBodyFromGRPCMessage(transport)
		require.Equal(t, initFrom, to)
	})

	t.Run("chunk non empty", func(t *testing.T) {
		wire, err := chunkFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.GetResponseBodyFromGRPCMessage(transport)
		require.Equal(t, chunkFrom, to)
	})

	t.Run("split info non empty", func(t *testing.T) {
		wire, err := splitInfoFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.GetResponseBodyFromGRPCMessage(transport)
		require.Equal(t, splitInfoFrom, to)
	})
}

func TestPutRequestBody_StableMarshal(t *testing.T) {
	initFrom := generatePutRequestBody(true)
	chunkFrom := generatePutRequestBody(false)
	transport := new(grpc.PutRequest_Body)

	t.Run("init non empty", func(t *testing.T) {
		wire, err := initFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.PutRequestBodyFromGRPCMessage(transport)
		require.Equal(t, initFrom, to)
	})

	t.Run("chunk non empty", func(t *testing.T) {
		wire, err := chunkFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
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

		err = goproto.Unmarshal(wire, transport)
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

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.DeleteRequestBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestDeleteResponseBody_StableMarshal(t *testing.T) {
	from := generateDeleteResponseBody("CID", "OID")
	transport := new(grpc.DeleteResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
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

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.HeadRequestBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestHeadResponseBody_StableMarshal(t *testing.T) {
	shortFrom := generateHeadResponseBody(0)
	fullFrom := generateHeadResponseBody(1)
	splitInfoFrom := generateHeadResponseBody(2)
	transport := new(grpc.HeadResponse_Body)

	t.Run("short header non empty", func(t *testing.T) {
		wire, err := shortFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.HeadResponseBodyFromGRPCMessage(transport)
		require.Equal(t, shortFrom, to)
	})

	t.Run("full header non empty", func(t *testing.T) {
		wire, err := fullFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.HeadResponseBodyFromGRPCMessage(transport)
		require.Equal(t, fullFrom, to)
	})

	t.Run("split info non empty", func(t *testing.T) {
		wire, err := splitInfoFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.HeadResponseBodyFromGRPCMessage(transport)
		require.Equal(t, splitInfoFrom, to)
	})
}

func TestSearchRequestBody_StableMarshal(t *testing.T) {
	from := generateSearchRequestBody(10, "Container ID")
	transport := new(grpc.SearchRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
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

		err = goproto.Unmarshal(wire, transport)
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

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.GetRangeRequestBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestGetRangeResponseBody_StableMarshal(t *testing.T) {
	dataFrom := generateRangeResponseBody("some data", true)
	splitInfoFrom := generateRangeResponseBody("some data", false)
	transport := new(grpc.GetRangeResponse_Body)

	t.Run("data non empty", func(t *testing.T) {
		wire, err := dataFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.GetRangeResponseBodyFromGRPCMessage(transport)
		require.Equal(t, dataFrom, to)
	})

	t.Run("split info non empty", func(t *testing.T) {
		wire, err := splitInfoFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.GetRangeResponseBodyFromGRPCMessage(transport)
		require.Equal(t, splitInfoFrom, to)
	})
}

func TestGetRangeHashRequestBody_StableMarshal(t *testing.T) {
	from := generateRangeHashRequestBody("Container ID", "Object ID", 5)
	transport := new(grpc.GetRangeHashRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		err = goproto.Unmarshal(wire, transport)
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

		err = goproto.Unmarshal(wire, transport)
		require.NoError(t, err)

		to := object.GetRangeHashResponseBodyFromGRPCMessage(transport)
		require.Equal(t, from, to)
	})
}

func TestHeaderWithSignature_StableMarshal(t *testing.T) {
	from := generateHeaderWithSignature()

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		to := new(object.HeaderWithSignature)
		require.NoError(t, to.Unmarshal(wire))

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
	hdr.SetPayloadHash(generateChecksum("payload hash"))
	hdr.SetHomomorphicHash(generateChecksum("homomorphic hash"))

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
	split.SetSplitID([]byte("UUIDv4"))

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
	hdr.SetSessionToken(generateSessionToken(strconv.Itoa(int(ln))))

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

func generateGetResponseBody(i int) *object.GetResponseBody {
	resp := new(object.GetResponseBody)
	var part object.GetObjectPart

	switch i {
	case 0:
		init := new(object.GetObjectPartInit)
		init.SetObjectID(generateObjectID("Object ID"))
		init.SetSignature(generateSignature("Key", "Signature"))
		init.SetHeader(generateHeader(10))
		part = init
	case 1:
		chunk := new(object.GetObjectPartChunk)
		chunk.SetChunk([]byte("Some data chunk"))
		part = chunk
	default:
		part = generateSplitInfo()
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

func generateDeleteResponseBody(cid, oid string) *object.DeleteResponseBody {
	resp := new(object.DeleteResponseBody)
	resp.SetTombstone(generateAddress(cid, oid))

	return resp
}

func generateHeadRequestBody(cid, oid string) *object.HeadRequestBody {
	req := new(object.HeadRequestBody)
	req.SetAddress(generateAddress(cid, oid))
	req.SetRaw(true)
	req.SetMainOnly(true)

	return req
}

func generateHeadResponseBody(flag int) *object.HeadResponseBody {
	req := new(object.HeadResponseBody)
	var part object.GetHeaderPart

	switch flag {
	case 0:
		part = generateShortHeader("short id")
	case 1:
		part = generateHeaderWithSignature()
	default:
		part = generateSplitInfo()
	}

	req.SetHeaderPart(part)

	return req
}

func generateHeaderWithSignature() *object.HeaderWithSignature {
	hdrWithSig := new(object.HeaderWithSignature)
	hdrWithSig.SetHeader(generateHeader(30))
	hdrWithSig.SetSignature(generateSignature("sig", "key"))

	return hdrWithSig
}

func generateFilter(k, v string) *object.SearchFilter {
	f := new(object.SearchFilter)
	f.SetKey(k)
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
	req.SetRaw(true)

	return req
}

func generateRangeResponseBody(data string, flag bool) *object.GetRangeResponseBody {
	resp := new(object.GetRangeResponseBody)

	if flag {
		p := new(object.GetRangePartChunk)
		p.SetChunk([]byte(data))
		resp.SetRangePart(p)
	} else {
		resp.SetRangePart(generateSplitInfo())
	}

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
		list[i] = []byte("Some homomorphic hash data" + strconv.Itoa(n))
	}

	resp.SetType(refs.TillichZemor)
	resp.SetHashList(list)

	return resp
}

func TestObject_StableUnmarshal(t *testing.T) {
	obj := generateObject("some data")

	data, err := obj.StableMarshal(nil)
	require.NoError(t, err)

	obj2 := new(object.Object)
	require.NoError(t, obj2.StableUnmarshal(data))

	require.Equal(t, obj, obj2)
}

func generateSplitInfo() *object.SplitInfo {
	splitInfo := new(object.SplitInfo)
	splitInfo.SetSplitID([]byte("splitID"))
	splitInfo.SetLastPart(generateObjectID("Right ID"))
	splitInfo.SetLink(generateObjectID("Link ID"))

	return splitInfo
}
