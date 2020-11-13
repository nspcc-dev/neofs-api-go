package eacl_test

import (
	"crypto/sha256"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/acl/eacl"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

// example how to create eACL tables in applications
func example() {
	record := eacl.CreateRecord(eacl.ActionDeny, eacl.OperationPut)
	record.AddFilter(eacl.HeaderFromObject, eacl.MatchStringEqual, "filename", "cat.jpg")
	record.AddTarget(eacl.RoleOthers, test.DecodeKey(1).PublicKey, test.DecodeKey(2).PublicKey)

	var cid container.ID
	cid.SetSHA256(sha256.Sum256([]byte("container id")))

	table := eacl.CreateTable(cid)
	table.AddRecord(record)
}

func TestTable(t *testing.T) {
	var (
		v   pkg.Version
		cid container.ID
	)

	sha := sha256.Sum256([]byte("container id"))
	cid.SetSHA256(sha)

	v.SetMajor(3)
	v.SetMinor(2)

	table := eacl.NewTable()
	table.SetVersion(v)
	table.SetCID(&cid)
	table.AddRecord(eacl.CreateRecord(eacl.ActionAllow, eacl.OperationPut))

	v2 := table.ToV2()
	require.NotNil(t, v2)
	require.Equal(t, uint32(3), v2.GetVersion().GetMajor())
	require.Equal(t, uint32(2), v2.GetVersion().GetMinor())
	require.Equal(t, sha[:], v2.GetContainerID().GetValue())
	require.Len(t, v2.GetRecords(), 1)

	newTable := eacl.NewTableFromV2(v2)
	require.Equal(t, table, newTable)

	t.Run("new from nil v2 table", func(t *testing.T) {
		require.Equal(t, new(eacl.Table), eacl.NewTableFromV2(nil))
	})

	t.Run("create table", func(t *testing.T) {
		var cid = new(container.ID)
		cid.SetSHA256(sha256.Sum256([]byte("container id")))

		table := eacl.CreateTable(*cid)
		require.Equal(t, cid, table.CID())
		require.Equal(t, *pkg.SDKVersion(), table.Version())
	})
}

func TestTable_AddRecord(t *testing.T) {
	records := []eacl.Record{
		*eacl.CreateRecord(eacl.ActionDeny, eacl.OperationDelete),
		*eacl.CreateRecord(eacl.ActionAllow, eacl.OperationPut),
	}

	table := eacl.NewTable()
	for _, record := range records {
		table.AddRecord(&record)
	}

	require.Equal(t, records, table.Records())
}

func TestRecordEncoding(t *testing.T) {
	tab := eacl.NewTable()
	tab.AddRecord(
		eacl.CreateRecord(eacl.ActionDeny, eacl.OperationHead),
	)

	t.Run("binary", func(t *testing.T) {
		data, err := tab.Marshal()
		require.NoError(t, err)

		tab2 := eacl.NewTable()
		require.NoError(t, tab2.Unmarshal(data))

		require.Equal(t, tab, tab2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := tab.MarshalJSON()
		require.NoError(t, err)

		r2 := eacl.NewTable()
		require.NoError(t, r2.UnmarshalJSON(data))

		require.Equal(t, tab, r2)
	})
}
