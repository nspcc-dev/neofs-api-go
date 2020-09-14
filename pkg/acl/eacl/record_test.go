package eacl

import (
	"crypto/ecdsa"
	"testing"

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
	record.AddTarget(RoleSystem)

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

func TestRecord_AddTarget(t *testing.T) {
	targets := []Target{
		{
			role: RoleUnknown,
			keys: []ecdsa.PublicKey{test.DecodeKey(1).PublicKey},
		},
		{
			role: RoleSystem,
			keys: []ecdsa.PublicKey{},
		},
	}

	r := NewRecord()
	for _, target := range targets {
		r.AddTarget(target.Role(), target.Keys()...)
	}

	require.Equal(t, targets, r.Targets())
}

func TestRecord_AddFilter(t *testing.T) {
	filters := []Filter{
		{
			from:    HeaderFromObject,
			name:    HdrObjSysNameCID,
			matcher: MatchStringEqual,
			value:   "ContainerID",
		},
		{
			from:    HeaderFromRequest,
			name:    "X-Header-Name",
			matcher: MatchStringNotEqual,
			value:   "X-Header-Value",
		},
	}

	r := NewRecord()
	for _, filter := range filters {
		r.AddFilter(filter.From(), filter.Matcher(), filter.Name(), filter.Value())
	}

	require.Equal(t, filters, r.Filters())
}
