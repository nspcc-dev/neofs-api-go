package eacl

import (
	"testing"

	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
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
