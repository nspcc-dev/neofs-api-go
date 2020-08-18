package object_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/object"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/object/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
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

func generateSignature(k, v string) *service.Signature {
	sig := new(service.Signature)
	sig.SetKey([]byte(k))
	sig.SetSign([]byte(v))

	return sig
}

func generateVersion(maj, min uint32) *service.Version {
	version := new(service.Version)
	version.SetMajor(maj)
	version.SetMinor(min)

	return version
}

func generateSessionToken(id string) *service.SessionToken {
	lifetime := new(service.TokenLifetime)
	lifetime.SetExp(1)
	lifetime.SetNbf(2)
	lifetime.SetIat(3)

	addr := new(refs.Address)
	addr.SetContainerID(generateContainerID("Container ID"))
	addr.SetObjectID(generateObjectID("Object ID"))

	objectCtx := new(service.ObjectSessionContext)
	objectCtx.SetVerb(service.ObjectVerbPut)
	objectCtx.SetAddress(addr)

	tokenBody := new(service.SessionTokenBody)
	tokenBody.SetID([]byte(id))
	tokenBody.SetOwnerID(generateOwner("Owner ID"))
	tokenBody.SetSessionKey([]byte(id))
	tokenBody.SetLifetime(lifetime)
	tokenBody.SetContext(objectCtx)

	sessionToken := new(service.SessionToken)
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
	hdr.SetHomomorphicHash([]byte("Homomorphic Hash"))
	hdr.SetObjectType(object.TypeRegular)
	hdr.SetPayloadHash([]byte("Payload Hash"))
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
