package eacl

import (
	"crypto/ecdsa"
	"fmt"
	"testing"

	cidtest "github.com/nspcc-dev/neofs-api-go/pkg/container/id/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	objecttest "github.com/nspcc-dev/neofs-api-go/pkg/object/test"
	ownertest "github.com/nspcc-dev/neofs-api-go/pkg/owner/test"
	refstest "github.com/nspcc-dev/neofs-api-go/pkg/test"
	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestRecord(t *testing.T) {
	record := NewRecord()
	record.SetOperation(OperationRange)
	record.SetAction(ActionAllow)
	record.AddFilter(HeaderFromRequest, MatchStringEqual, "A", "B")
	record.AddFilter(HeaderFromRequest, MatchStringNotEqual, "C", "D")

	target := NewTarget()
	target.SetRole(RoleSystem)
	AddRecordTarget(record, target)

	v2 := record.ToV2()
	require.NotNil(t, v2)
	require.Equal(t, v2acl.OperationRange, v2.GetOperation())
	require.Equal(t, v2acl.ActionAllow, v2.GetAction())
	require.Len(t, v2.GetFilters(), len(record.Filters()))
	require.Len(t, v2.GetTargets(), len(record.Targets()))

	newRecord := NewRecordFromV2(v2)
	require.Equal(t, record, newRecord)

	t.Run("create record", func(t *testing.T) {
		record := CreateRecord(ActionAllow, OperationGet)
		require.Equal(t, ActionAllow, record.Action())
		require.Equal(t, OperationGet, record.Operation())
	})

	t.Run("new from nil v2 record", func(t *testing.T) {
		require.Equal(t, new(Record), NewRecordFromV2(nil))
	})
}

func TestAddFormedTarget(t *testing.T) {
	items := []struct {
		role Role
		keys []ecdsa.PublicKey
	}{
		{
			role: RoleUnknown,
			keys: []ecdsa.PublicKey{test.DecodeKey(1).PublicKey},
		},
		{
			role: RoleSystem,
			keys: []ecdsa.PublicKey{},
		},
	}

	targets := make([]*Target, 0, len(items))

	r := NewRecord()

	for _, item := range items {
		tgt := NewTarget()
		tgt.SetRole(item.role)
		SetTargetECDSAKeys(tgt, ecdsaKeysToPtrs(item.keys)...)

		targets = append(targets, tgt)

		AddFormedTarget(r, item.role, item.keys...)
	}

	tgts := r.Targets()
	require.Len(t, tgts, len(targets))

	for _, tgt := range targets {
		require.Contains(t, tgts, tgt)
	}
}

func TestRecord_AddFilter(t *testing.T) {
	filters := []*Filter{
		newObjectFilter(MatchStringEqual, "some name", "ContainerID"),
		newObjectFilter(MatchStringNotEqual, "X-Header-Name", "X-Header-Value"),
	}

	r := NewRecord()
	for _, filter := range filters {
		r.AddFilter(filter.From(), filter.Matcher(), filter.Key(), filter.Value())
	}

	require.Equal(t, filters, r.Filters())
}

func TestRecordEncoding(t *testing.T) {
	r := NewRecord()
	r.SetOperation(OperationHead)
	r.SetAction(ActionDeny)
	r.AddObjectAttributeFilter(MatchStringEqual, "key", "value")
	AddFormedTarget(r, RoleSystem, test.DecodeKey(-1).PublicKey)

	t.Run("binary", func(t *testing.T) {
		data, err := r.Marshal()
		require.NoError(t, err)

		r2 := NewRecord()
		require.NoError(t, r2.Unmarshal(data))

		require.Equal(t, r, r2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := r.MarshalJSON()
		require.NoError(t, err)

		r2 := NewRecord()
		require.NoError(t, r2.UnmarshalJSON(data))

		require.Equal(t, r, r2)
	})
}

func TestRecord_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *Record

		require.Nil(t, x.ToV2())
	})

	t.Run("default values", func(t *testing.T) {
		record := NewRecord()

		// check initial values
		require.Equal(t, OperationUnknown, record.Operation())
		require.Equal(t, ActionUnknown, record.Action())
		require.Nil(t, record.Targets())
		require.Nil(t, record.Filters())

		// convert to v2 message
		recordV2 := record.ToV2()

		require.Equal(t, v2acl.OperationUnknown, recordV2.GetOperation())
		require.Equal(t, v2acl.ActionUnknown, recordV2.GetAction())
		require.Nil(t, recordV2.GetTargets())
		require.Nil(t, recordV2.GetFilters())
	})
}

func TestReservedRecords(t *testing.T) {
	var (
		v       = refstest.Version()
		oid     = objecttest.ID()
		cid     = cidtest.Generate()
		ownerid = ownertest.Generate()
		h       = refstest.Checksum()
		typ     = new(object.Type)
	)

	testSuit := []struct {
		f     func(r *Record)
		key   string
		value string
	}{
		{
			f:     func(r *Record) { r.AddObjectAttributeFilter(MatchStringEqual, "foo", "bar") },
			key:   "foo",
			value: "bar",
		},
		{
			f:     func(r *Record) { r.AddObjectVersionFilter(MatchStringEqual, v) },
			key:   v2acl.FilterObjectVersion,
			value: v.String(),
		},
		{
			f:     func(r *Record) { r.AddObjectIDFilter(MatchStringEqual, oid) },
			key:   v2acl.FilterObjectID,
			value: oid.String(),
		},
		{
			f:     func(r *Record) { r.AddObjectContainerIDFilter(MatchStringEqual, cid) },
			key:   v2acl.FilterObjectContainerID,
			value: cid.String(),
		},
		{
			f:     func(r *Record) { r.AddObjectOwnerIDFilter(MatchStringEqual, ownerid) },
			key:   v2acl.FilterObjectOwnerID,
			value: ownerid.String(),
		},
		{
			f:     func(r *Record) { r.AddObjectCreationEpoch(MatchStringEqual, 100) },
			key:   v2acl.FilterObjectCreationEpoch,
			value: "100",
		},
		{
			f:     func(r *Record) { r.AddObjectPayloadLengthFilter(MatchStringEqual, 5000) },
			key:   v2acl.FilterObjectPayloadLength,
			value: "5000",
		},
		{
			f:     func(r *Record) { r.AddObjectPayloadHashFilter(MatchStringEqual, h) },
			key:   v2acl.FilterObjectPayloadHash,
			value: h.String(),
		},
		{
			f:     func(r *Record) { r.AddObjectHomomorphicHashFilter(MatchStringEqual, h) },
			key:   v2acl.FilterObjectHomomorphicHash,
			value: h.String(),
		},
		{
			f: func(r *Record) {
				require.True(t, typ.FromString("REGULAR"))
				r.AddObjectTypeFilter(MatchStringEqual, *typ)
			},
			key:   v2acl.FilterObjectType,
			value: "REGULAR",
		},
		{
			f: func(r *Record) {
				require.True(t, typ.FromString("TOMBSTONE"))
				r.AddObjectTypeFilter(MatchStringEqual, *typ)
			},
			key:   v2acl.FilterObjectType,
			value: "TOMBSTONE",
		},
		{
			f: func(r *Record) {
				require.True(t, typ.FromString("STORAGE_GROUP"))
				r.AddObjectTypeFilter(MatchStringEqual, *typ)
			},
			key:   v2acl.FilterObjectType,
			value: "STORAGE_GROUP",
		},
	}

	for n, testCase := range testSuit {
		desc := fmt.Sprintf("case #%d", n)
		record := NewRecord()
		testCase.f(record)
		require.Len(t, record.Filters(), 1, desc)
		f := record.Filters()[0]
		require.Equal(t, f.Key(), testCase.key, desc)
		require.Equal(t, f.Value(), testCase.value, desc)
	}
}
