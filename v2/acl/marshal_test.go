package acl_test

import (
	"fmt"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func generateTarget(u acl.Role, k int) *acl.Target {
	target := new(acl.Target)
	target.SetRole(u)

	keys := make([][]byte, k)

	for i := 0; i < k; i++ {
		s := fmt.Sprintf("Public Key %d", i+1)
		keys[i] = []byte(s)
	}

	return target
}

func generateFilter(t acl.HeaderType, k, v string) *acl.HeaderFilter {
	filter := new(acl.HeaderFilter)
	filter.SetHeaderType(t)
	filter.SetMatchType(acl.MatchTypeStringEqual)
	filter.SetName(k)
	filter.SetValue(v)

	return filter
}

func generateRecord(another bool) *acl.Record {
	record := new(acl.Record)

	switch another {
	case true:
		t1 := generateTarget(acl.RoleUser, 2)
		f1 := generateFilter(acl.HeaderTypeObject, "OID", "ObjectID Value")

		record.SetOperation(acl.OperationHead)
		record.SetAction(acl.ActionDeny)
		record.SetTargets([]*acl.Target{t1})
		record.SetFilters([]*acl.HeaderFilter{f1})
	default:
		t1 := generateTarget(acl.RoleUser, 2)
		t2 := generateTarget(acl.RoleSystem, 0)
		f1 := generateFilter(acl.HeaderTypeObject, "CID", "Container ID Value")
		f2 := generateFilter(acl.HeaderTypeRequest, "X-Header-Key", "X-Header-Value")

		record.SetOperation(acl.OperationPut)
		record.SetAction(acl.ActionAllow)
		record.SetTargets([]*acl.Target{t1, t2})
		record.SetFilters([]*acl.HeaderFilter{f1, f2})
	}

	return record
}

func generateEACL() *acl.Table {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte("Container ID"))

	ver := new(refs.Version)
	ver.SetMajor(2)
	ver.SetMinor(3)

	table := new(acl.Table)
	table.SetVersion(ver)
	table.SetContainerID(cid)
	table.SetRecords([]*acl.Record{generateRecord(true)})

	return table
}

func generateSignature(k, v string) *refs.Signature {
	sig := new(refs.Signature)
	sig.SetKey([]byte(k))
	sig.SetSign([]byte(v))

	return sig
}

func generateLifetime(exp, nbf, iat uint64) *acl.TokenLifetime {
	lifetime := new(acl.TokenLifetime)
	lifetime.SetExp(exp)
	lifetime.SetNbf(nbf)
	lifetime.SetIat(iat)

	return lifetime
}

func generateBearerTokenBody(id string) *acl.BearerTokenBody {
	owner := new(refs.OwnerID)
	owner.SetValue([]byte(id))

	tokenBody := new(acl.BearerTokenBody)
	tokenBody.SetOwnerID(owner)
	tokenBody.SetLifetime(generateLifetime(1, 2, 3))
	tokenBody.SetEACL(generateEACL())

	return tokenBody
}

func generateBearerToken(id string) *acl.BearerToken {
	bearerToken := new(acl.BearerToken)
	bearerToken.SetBody(generateBearerTokenBody(id))
	bearerToken.SetSignature(generateSignature("id", id))

	return bearerToken
}

func TestHeaderFilter_StableMarshal(t *testing.T) {
	filterFrom := generateFilter(acl.HeaderTypeObject, "CID", "Container ID Value")
	transport := new(grpc.EACLRecord_Filter)

	t.Run("non empty", func(t *testing.T) {
		filterFrom.SetHeaderType(acl.HeaderTypeObject)
		filterFrom.SetMatchType(acl.MatchTypeStringEqual)
		filterFrom.SetName("Hello")
		filterFrom.SetValue("World")

		wire, err := filterFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		filterTo := acl.HeaderFilterFromGRPCMessage(transport)
		require.Equal(t, filterFrom, filterTo)
	})
}

func TestTargetInfo_StableMarshal(t *testing.T) {
	targetFrom := generateTarget(acl.RoleUser, 2)
	transport := new(grpc.EACLRecord_Target)

	t.Run("non empty", func(t *testing.T) {
		targetFrom.SetRole(acl.RoleUser)
		targetFrom.SetKeys([][]byte{
			[]byte("Public Key 1"),
			[]byte("Public Key 2"),
		})

		wire, err := targetFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		targetTo := acl.TargetInfoFromGRPCMessage(transport)
		require.Equal(t, targetFrom, targetTo)
	})
}

func TestRecord_StableMarshal(t *testing.T) {
	recordFrom := generateRecord(false)
	transport := new(grpc.EACLRecord)

	t.Run("non empty", func(t *testing.T) {
		wire, err := recordFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		recordTo := acl.RecordFromGRPCMessage(transport)
		require.Equal(t, recordFrom, recordTo)
	})
}

func TestTable_StableMarshal(t *testing.T) {
	tableFrom := new(acl.Table)
	transport := new(grpc.EACLTable)

	t.Run("non empty", func(t *testing.T) {
		cid := new(refs.ContainerID)
		cid.SetValue([]byte("Container ID"))

		ver := new(refs.Version)
		ver.SetMajor(2)
		ver.SetMinor(3)

		r1 := generateRecord(false)
		r2 := generateRecord(true)

		tableFrom.SetVersion(ver)
		tableFrom.SetContainerID(cid)
		tableFrom.SetRecords([]*acl.Record{r1, r2})

		wire, err := tableFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		tableTo := acl.TableFromGRPCMessage(transport)
		require.Equal(t, tableFrom, tableTo)
	})
}

func TestTokenLifetime_StableMarshal(t *testing.T) {
	lifetimeFrom := generateLifetime(10, 20, 30)
	transport := new(grpc.BearerToken_Body_TokenLifetime)

	t.Run("non empty", func(t *testing.T) {
		wire, err := lifetimeFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		lifetimeTo := acl.TokenLifetimeFromGRPCMessage(transport)
		require.Equal(t, lifetimeFrom, lifetimeTo)
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

		bearerTokenBodyTo := acl.BearerTokenBodyFromGRPCMessage(transport)
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

		bearerTokenTo := acl.BearerTokenFromGRPCMessage(transport)
		require.Equal(t, bearerTokenFrom, bearerTokenTo)
	})
}
