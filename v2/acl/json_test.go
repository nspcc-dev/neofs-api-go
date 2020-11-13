package acl_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/stretchr/testify/require"
)

func TestBearerTokenJSON(t *testing.T) {
	exp := generateBearerToken("token")

	t.Run("non empty", func(t *testing.T) {
		data, err := acl.BearerTokenToJSON(exp)
		require.NoError(t, err)

		got, err := acl.BearerTokenFromJSON(data)
		require.NoError(t, err)

		require.Equal(t, exp, got)
	})
}

func TestFilterJSON(t *testing.T) {
	f := generateFilter(acl.HeaderTypeObject, "key", "value")

	data, err := f.MarshalJSON()
	require.NoError(t, err)

	f2 := new(acl.HeaderFilter)
	require.NoError(t, f2.UnmarshalJSON(data))

	require.Equal(t, f, f2)
}

func TestTargetJSON(t *testing.T) {
	tar := generateTarget(acl.RoleSystem, 3)

	data, err := tar.MarshalJSON()
	require.NoError(t, err)

	tar2 := new(acl.Target)
	require.NoError(t, tar2.UnmarshalJSON(data))

	require.Equal(t, tar, tar2)
}

func TestTable_MarshalJSON(t *testing.T) {
	tab := new(acl.Table)
	tab.SetRecords([]*acl.Record{generateRecord(false)})

	data, err := tab.MarshalJSON()
	require.NoError(t, err)

	tab2 := new(acl.Table)
	require.NoError(t, tab2.UnmarshalJSON(data))

	require.Equal(t, tab, tab2)
}
