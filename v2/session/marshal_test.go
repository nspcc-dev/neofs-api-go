package session_test

import (
	"fmt"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
	"github.com/stretchr/testify/require"
)

func TestCreateRequestBody_StableMarshal(t *testing.T) {
	requestFrom := generateCreateSessionRequestBody("Owner ID")
	transport := new(grpc.CreateRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := requestFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		requestTo := session.CreateRequestBodyFromGRPCMessage(transport)
		require.Equal(t, requestFrom, requestTo)
	})
}

func TestCreateResponseBody_StableMarshal(t *testing.T) {
	responseFrom := generateCreateSessionResponseBody("ID", "Session Public Key")
	transport := new(grpc.CreateResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := responseFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		responseTo := session.CreateResponseBodyFromGRPCMessage(transport)
		require.Equal(t, responseFrom, responseTo)
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

		xheaderTo := session.XHeaderFromGRPCMessage(transport)
		require.Equal(t, xheaderFrom, xheaderTo)
	})
}

func TestTokenLifetime_StableMarshal(t *testing.T) {
	lifetimeFrom := generateLifetime(10, 20, 30)
	transport := new(grpc.SessionToken_Body_TokenLifetime)

	t.Run("non empty", func(t *testing.T) {
		wire, err := lifetimeFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		lifetimeTo := session.TokenLifetimeFromGRPCMessage(transport)
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

		objectCtxTo := session.ObjectSessionContextFromGRPCMessage(transport)
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

		sessionTokenBodyTo := session.SessionTokenBodyFromGRPCMessage(transport)
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

		sessionTokenTo := session.SessionTokenFromGRPCMessage(transport)
		require.Equal(t, sessionTokenFrom, sessionTokenTo)
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

		metaHeaderTo := session.RequestMetaHeaderFromGRPCMessage(transport)
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

		verifHeaderTo := session.RequestVerificationHeaderFromGRPCMessage(transport)
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

		metaHeaderTo := session.ResponseMetaHeaderFromGRPCMessage(transport)
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

		verifHeaderTo := session.ResponseVerificationHeaderFromGRPCMessage(transport)
		require.Equal(t, verifHeaderFrom, verifHeaderTo)
	})
}

func generateCreateSessionRequestBody(id string) *session.CreateRequestBody {
	lifetime := new(session.TokenLifetime)
	lifetime.SetIat(1)
	lifetime.SetNbf(2)
	lifetime.SetExp(3)

	owner := new(refs.OwnerID)
	owner.SetValue([]byte(id))

	s := new(session.CreateRequestBody)
	s.SetOwnerID(owner)
	s.SetExpiration(10)

	return s
}

func generateCreateSessionResponseBody(id, key string) *session.CreateResponseBody {
	s := new(session.CreateResponseBody)
	s.SetID([]byte(id))
	s.SetSessionKey([]byte(key))

	return s
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

func generateXHeader(k, v string) *session.XHeader {
	xheader := new(session.XHeader)
	xheader.SetKey(k)
	xheader.SetValue(v)

	return xheader
}

func generateLifetime(exp, nbf, iat uint64) *session.TokenLifetime {
	lifetime := new(session.TokenLifetime)
	lifetime.SetExp(exp)
	lifetime.SetNbf(nbf)
	lifetime.SetIat(iat)

	return lifetime
}

func generateBearerLifetime(exp, nbf, iat uint64) *acl.TokenLifetime {
	lifetime := new(acl.TokenLifetime)
	lifetime.SetExp(exp)
	lifetime.SetNbf(nbf)
	lifetime.SetIat(iat)

	return lifetime
}

func generateObjectCtx(id string) *session.ObjectSessionContext {
	objectCtx := new(session.ObjectSessionContext)
	objectCtx.SetVerb(session.ObjectVerbPut)

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
	target.SetRole(acl.RoleUser)

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

func generateSessionTokenBody(id string) *session.SessionTokenBody {
	owner := new(refs.OwnerID)
	owner.SetValue([]byte("Owner ID"))

	tokenBody := new(session.SessionTokenBody)
	tokenBody.SetID([]byte(id))
	tokenBody.SetOwnerID(owner)
	tokenBody.SetSessionKey([]byte(id))
	tokenBody.SetLifetime(generateLifetime(1, 2, 3))
	tokenBody.SetContext(generateObjectCtx(id))

	return tokenBody
}

func generateSessionToken(id string) *session.SessionToken {
	sessionToken := new(session.SessionToken)
	sessionToken.SetBody(generateSessionTokenBody(id))
	sessionToken.SetSignature(generateSignature("id", id))

	return sessionToken
}

func generateBearerTokenBody(id string) *acl.BearerTokenBody {
	owner := new(refs.OwnerID)
	owner.SetValue([]byte(id))

	tokenBody := new(acl.BearerTokenBody)
	tokenBody.SetOwnerID(owner)
	tokenBody.SetLifetime(generateBearerLifetime(1, 2, 3))
	tokenBody.SetEACL(generateEACL(10, "id", id))

	return tokenBody
}

func generateBearerToken(id string) *acl.BearerToken {
	bearerToken := new(acl.BearerToken)
	bearerToken.SetBody(generateBearerTokenBody(id))
	bearerToken.SetSignature(generateSignature("id", id))

	return bearerToken
}

func generateRequestMetaHeader(n int, b, s string) *session.RequestMetaHeader {
	reqMetaHeader := new(session.RequestMetaHeader)
	reqMetaHeader.SetVersion(generateVersion(2, 0))
	reqMetaHeader.SetEpoch(uint64(n))
	reqMetaHeader.SetTTL(uint32(n))
	reqMetaHeader.SetXHeaders([]*session.XHeader{
		generateXHeader("key-one", "val-one"),
		generateXHeader("key-two", "val-two"),
	})
	reqMetaHeader.SetBearerToken(generateBearerToken(b))
	reqMetaHeader.SetSessionToken(generateSessionToken(s))

	return reqMetaHeader
}

func generateRequestVerificationHeader(k, v string) *session.RequestVerificationHeader {
	reqVerifHeader := new(session.RequestVerificationHeader)
	reqVerifHeader.SetBodySignature(generateSignature(k+"body", v+"body"))
	reqVerifHeader.SetMetaSignature(generateSignature(k+"meta", v+"meta"))
	reqVerifHeader.SetOriginSignature(generateSignature(k+"orig", v+"orig"))

	return reqVerifHeader
}

func generateResponseMetaHeader(n int) *session.ResponseMetaHeader {
	respMetaHeader := new(session.ResponseMetaHeader)
	respMetaHeader.SetVersion(generateVersion(2, 0))
	respMetaHeader.SetEpoch(uint64(n))
	respMetaHeader.SetTTL(uint32(n))
	respMetaHeader.SetXHeaders([]*session.XHeader{
		generateXHeader("key-one", "val-one"),
		generateXHeader("key-two", "val-two"),
	})

	return respMetaHeader
}

func generateResponseVerificationHeader(k, v string) *session.ResponseVerificationHeader {
	respVerifHeader := new(session.ResponseVerificationHeader)
	respVerifHeader.SetBodySignature(generateSignature(k+"body", v+"body"))
	respVerifHeader.SetMetaSignature(generateSignature(k+"meta", v+"meta"))
	respVerifHeader.SetOriginSignature(generateSignature(k+"orig", v+"orig"))

	return respVerifHeader
}
