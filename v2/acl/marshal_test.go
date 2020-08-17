package acl_test

import (
	"fmt"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func generateTarget(u acl.Target, k int) *acl.TargetInfo {
	target := new(acl.TargetInfo)
	target.SetTarget(u)

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
		t1 := generateTarget(acl.TargetUser, 2)
		f1 := generateFilter(acl.HeaderTypeObject, "OID", "ObjectID Value")

		record.SetOperation(acl.OperationHead)
		record.SetAction(acl.ActionDeny)
		record.SetTargets([]*acl.TargetInfo{t1})
		record.SetFilters([]*acl.HeaderFilter{f1})
	default:
		t1 := generateTarget(acl.TargetUser, 2)
		t2 := generateTarget(acl.TargetSystem, 0)
		f1 := generateFilter(acl.HeaderTypeObject, "CID", "Container ID Value")
		f2 := generateFilter(acl.HeaderTypeRequest, "X-Header-Key", "X-Header-Value")

		record.SetOperation(acl.OperationPut)
		record.SetAction(acl.ActionAllow)
		record.SetTargets([]*acl.TargetInfo{t1, t2})
		record.SetFilters([]*acl.HeaderFilter{f1, f2})
	}

	return record
}

func TestHeaderFilter_StableMarshal(t *testing.T) {
	filterFrom := generateFilter(acl.HeaderTypeObject, "CID", "Container ID Value")
	transport := new(grpc.EACLRecord_FilterInfo)

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
	targetFrom := generateTarget(acl.TargetUser, 2)
	transport := new(grpc.EACLRecord_TargetInfo)

	t.Run("non empty", func(t *testing.T) {
		targetFrom.SetTarget(acl.TargetUser)
		targetFrom.SetKeyList([][]byte{
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

		r1 := generateRecord(false)
		r2 := generateRecord(true)

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
