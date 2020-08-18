package service_test

import (
	"fmt"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/service/grpc"
	"github.com/stretchr/testify/require"
)

func TestSignature_StableMarshal(t *testing.T) {
	signatureFrom := generateSignature("Public Key", "Signature")
	transport := new(grpc.Signature)

	t.Run("non empty", func(t *testing.T) {
		wire, err := signatureFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		signatureTo := service.SignatureFromGRPCMessage(transport)
		require.Equal(t, signatureFrom, signatureTo)
	})
}

func TestVersion_StableMarshal(t *testing.T) {
	versionFrom := generateVersion(2, 0)
	transport := new(grpc.Version)

	t.Run("non empty", func(t *testing.T) {
		wire, err := versionFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		versionTo := service.VersionFromGRPCMessage(transport)
		require.Equal(t, versionFrom, versionTo)
	})
}

func TestXHeader_StableMarshal(t *testing.T) {
	xheaderFrom := generateXHeader("X-Header-Key", "X-Header-Value")
	transport := new(grpc.XHeader)

	t.Run("non empty", func(t *testing.T) {
		wire, err := xheaderFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		xheaderTo := service.XHeaderFromGRPCMessage(transport)
		require.Equal(t, xheaderFrom, xheaderTo)
	})
}

func TestTokenLifetime_StableMarshal(t *testing.T) {
	lifetimeFrom := generateLifetime(10, 20, 30)
	transport := new(grpc.TokenLifetime)

	t.Run("non empty", func(t *testing.T) {
		wire, err := lifetimeFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		lifetimeTo := service.TokenLifetimeFromGRPCMessage(transport)
		require.Equal(t, lifetimeFrom, lifetimeTo)
	})
}

func TestObjectSessionContext_StableMarshal(t *testing.T) {
	objectCtxFrom := generateObjectCtx("Object ID")
	transport := new(grpc.ObjectSessionContext)

	t.Run("non empty", func(t *testing.T) {
		wire, err := objectCtxFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		objectCtxTo := service.ObjectSessionContextFromGRPCMessage(transport)
		require.Equal(t, objectCtxFrom, objectCtxTo)
	})
}

func TestSessionTokenBody_StableMarshal(t *testing.T) {
	sessionTokenBodyFrom := generateSessionTokenBody("Session Token Body")
	transport := new(grpc.SessionToken_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := sessionTokenBodyFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		sessionTokenBodyTo := service.SessionTokenBodyFromGRPCMessage(transport)
		require.Equal(t, sessionTokenBodyFrom, sessionTokenBodyTo)
	})
}

func TestSessionToken_StableMarshal(t *testing.T) {
	sessionTokenFrom := generateSessionToken("Session Token")
	transport := new(grpc.SessionToken)

	t.Run("non empty", func(t *testing.T) {
		wire, err := sessionTokenFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		sessionTokenTo := service.SessionTokenFromGRPCMessage(transport)
		require.Equal(t, sessionTokenFrom, sessionTokenTo)
	})
}

func TestBearerTokenBody_StableMarshal(t *testing.T) {
	bearerTokenBodyFrom := generateBearerTokenBody("Bearer Token Body")
	transport := new(grpc.BearerToken_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := bearerTokenBodyFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		bearerTokenBodyTo := service.BearerTokenBodyFromGRPCMessage(transport)
		require.Equal(t, bearerTokenBodyFrom, bearerTokenBodyTo)
	})
}

func TestBearerToken_StableMarshal(t *testing.T) {
	bearerTokenFrom := generateBearerToken("Bearer Token")
	transport := new(grpc.BearerToken)

	t.Run("non empty", func(t *testing.T) {
		wire, err := bearerTokenFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		bearerTokenTo := service.BearerTokenFromGRPCMessage(transport)
		require.Equal(t, bearerTokenFrom, bearerTokenTo)
	})
}

func TestRequestMetaHeader_StableMarshal(t *testing.T) {
	metaHeaderOrigin := generateRequestMetaHeader(10, "Bearer One", "Session One")
	metaHeaderFrom := generateRequestMetaHeader(20, "Bearer Two", "Session Two")
	metaHeaderFrom.SetOrigin(metaHeaderOrigin)
	transport := new(grpc.RequestMetaHeader)

	t.Run("non empty", func(t *testing.T) {
		wire, err := metaHeaderFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		metaHeaderTo := service.RequestMetaHeaderFromGRPCMessage(transport)
		require.Equal(t, metaHeaderFrom, metaHeaderTo)
	})
}

func TestRequestVerificationHeader_StableMarshal(t *testing.T) {
	verifHeaderOrigin := generateRequestVerificationHeader("Key", "Inside")
	verifHeaderFrom := generateRequestVerificationHeader("Value", "Outside")
	verifHeaderFrom.SetOrigin(verifHeaderOrigin)
	transport := new(grpc.RequestVerificationHeader)

	t.Run("non empty", func(t *testing.T) {
		wire, err := verifHeaderFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		verifHeaderTo := service.RequestVerificationHeaderFromGRPCMessage(transport)
		require.Equal(t, verifHeaderFrom, verifHeaderTo)
	})
}

func TestResponseMetaHeader_StableMarshal(t *testing.T) {
	metaHeaderOrigin := generateResponseMetaHeader(10)
	metaHeaderFrom := generateResponseMetaHeader(20)
	metaHeaderFrom.SetOrigin(metaHeaderOrigin)
	transport := new(grpc.ResponseMetaHeader)

	t.Run("non empty", func(t *testing.T) {
		wire, err := metaHeaderFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		metaHeaderTo := service.ResponseMetaHeaderFromGRPCMessage(transport)
		require.Equal(t, metaHeaderFrom, metaHeaderTo)
	})
}

func TestResponseVerificationHeader_StableMarshal(t *testing.T) {
	verifHeaderOrigin := generateResponseVerificationHeader("Key", "Inside")
	verifHeaderFrom := generateResponseVerificationHeader("Value", "Outside")
	verifHeaderFrom.SetOrigin(verifHeaderOrigin)
	transport := new(grpc.ResponseVerificationHeader)

	t.Run("non empty", func(t *testing.T) {
		wire, err := verifHeaderFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		verifHeaderTo := service.ResponseVerificationHeaderFromGRPCMessage(transport)
		require.Equal(t, verifHeaderFrom, verifHeaderTo)
	})
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

func generateXHeader(k, v string) *service.XHeader {
	xheader := new(service.XHeader)
	xheader.SetKey(k)
	xheader.SetValue(v)

	return xheader
}

func generateLifetime(exp, nbf, iat uint64) *service.TokenLifetime {
	lifetime := new(service.TokenLifetime)
	lifetime.SetExp(exp)
	lifetime.SetNbf(nbf)
	lifetime.SetIat(iat)

	return lifetime
}

func generateObjectCtx(id string) *service.ObjectSessionContext {
	objectCtx := new(service.ObjectSessionContext)
	objectCtx.SetVerb(service.ObjectVerbPut)

	cid := new(refs.ContainerID)
	cid.SetValue([]byte("ContainerID"))

	oid := new(refs.ObjectID)
	oid.SetValue([]byte(id))

	addr := new(refs.Address)
	addr.SetContainerID(cid)
	addr.SetObjectID(oid)

	objectCtx.SetAddress(addr)

	return objectCtx
}

func generateEACL(n int, k, v string) *acl.Table {
	target := new(acl.TargetInfo)
	target.SetTarget(acl.TargetUser)

	keys := make([][]byte, n)

	for i := 0; i < n; i++ {
		s := fmt.Sprintf("Public Key %d", i+1)
		keys[i] = []byte(s)
	}

	filter := new(acl.HeaderFilter)
	filter.SetHeaderType(acl.HeaderTypeObject)
	filter.SetMatchType(acl.MatchTypeStringEqual)
	filter.SetName(k)
	filter.SetValue(v)

	record := new(acl.Record)
	record.SetOperation(acl.OperationHead)
	record.SetAction(acl.ActionDeny)
	record.SetTargets([]*acl.TargetInfo{target})
	record.SetFilters([]*acl.HeaderFilter{filter})

	table := new(acl.Table)
	cid := new(refs.ContainerID)
	cid.SetValue([]byte("Container ID"))

	table.SetContainerID(cid)
	table.SetRecords([]*acl.Record{record})

	return table
}

func generateSessionTokenBody(id string) *service.SessionTokenBody {
	owner := new(refs.OwnerID)
	owner.SetValue([]byte("Owner ID"))

	tokenBody := new(service.SessionTokenBody)
	tokenBody.SetID([]byte(id))
	tokenBody.SetOwnerID(owner)
	tokenBody.SetSessionKey([]byte(id))
	tokenBody.SetLifetime(generateLifetime(1, 2, 3))
	tokenBody.SetContext(generateObjectCtx(id))

	return tokenBody
}

func generateSessionToken(id string) *service.SessionToken {
	sessionToken := new(service.SessionToken)
	sessionToken.SetBody(generateSessionTokenBody(id))
	sessionToken.SetSignature(generateSignature("id", id))

	return sessionToken
}

func generateBearerTokenBody(id string) *service.BearerTokenBody {
	owner := new(refs.OwnerID)
	owner.SetValue([]byte(id))

	tokenBody := new(service.BearerTokenBody)
	tokenBody.SetOwnerID(owner)
	tokenBody.SetLifetime(generateLifetime(1, 2, 3))
	tokenBody.SetEACL(generateEACL(10, "id", id))

	return tokenBody
}

func generateBearerToken(id string) *service.BearerToken {
	bearerToken := new(service.BearerToken)
	bearerToken.SetBody(generateBearerTokenBody(id))
	bearerToken.SetSignature(generateSignature("id", id))

	return bearerToken
}

func generateRequestMetaHeader(n int, b, s string) *service.RequestMetaHeader {
	reqMetaHeader := new(service.RequestMetaHeader)
	reqMetaHeader.SetVersion(generateVersion(2, 0))
	reqMetaHeader.SetEpoch(uint64(n))
	reqMetaHeader.SetTTL(uint32(n))
	reqMetaHeader.SetXHeaders([]*service.XHeader{
		generateXHeader("key-one", "val-one"),
		generateXHeader("key-two", "val-two"),
	})
	reqMetaHeader.SetBearerToken(generateBearerToken(b))
	reqMetaHeader.SetSessionToken(generateSessionToken(s))

	return reqMetaHeader
}

func generateRequestVerificationHeader(k, v string) *service.RequestVerificationHeader {
	reqVerifHeader := new(service.RequestVerificationHeader)
	reqVerifHeader.SetBodySignature(generateSignature(k+"body", v+"body"))
	reqVerifHeader.SetMetaSignature(generateSignature(k+"meta", v+"meta"))
	reqVerifHeader.SetOriginSignature(generateSignature(k+"orig", v+"orig"))

	return reqVerifHeader
}

func generateResponseMetaHeader(n int) *service.ResponseMetaHeader {
	respMetaHeader := new(service.ResponseMetaHeader)
	respMetaHeader.SetVersion(generateVersion(2, 0))
	respMetaHeader.SetEpoch(uint64(n))
	respMetaHeader.SetTTL(uint32(n))
	respMetaHeader.SetXHeaders([]*service.XHeader{
		generateXHeader("key-one", "val-one"),
		generateXHeader("key-two", "val-two"),
	})

	return respMetaHeader
}

func generateResponseVerificationHeader(k, v string) *service.ResponseVerificationHeader {
	respVerifHeader := new(service.ResponseVerificationHeader)
	respVerifHeader.SetBodySignature(generateSignature(k+"body", v+"body"))
	respVerifHeader.SetMetaSignature(generateSignature(k+"meta", v+"meta"))
	respVerifHeader.SetOriginSignature(generateSignature(k+"orig", v+"orig"))

	return respVerifHeader
}
