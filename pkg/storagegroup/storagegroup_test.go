package storagegroup_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	objecttest "github.com/nspcc-dev/neofs-api-go/pkg/object/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/storagegroup"
	storagegrouptest "github.com/nspcc-dev/neofs-api-go/pkg/storagegroup/test"
	refstest "github.com/nspcc-dev/neofs-api-go/pkg/test"
	storagegroupV2 "github.com/nspcc-dev/neofs-api-go/v2/storagegroup"
	"github.com/stretchr/testify/require"
)

func TestStorageGroup(t *testing.T) {
	sg := storagegroup.New()

	sz := uint64(13)
	sg.SetValidationDataSize(sz)
	require.Equal(t, sz, sg.ValidationDataSize())

	cs := refstest.Checksum()
	sg.SetValidationDataHash(cs)
	require.Equal(t, cs, sg.ValidationDataHash())

	exp := uint64(33)
	sg.SetExpirationEpoch(exp)
	require.Equal(t, exp, sg.ExpirationEpoch())

	members := []*object.ID{objecttest.ID(), objecttest.ID()}
	sg.SetMembers(members)
	require.Equal(t, members, sg.Members())
}

func TestStorageGroupEncoding(t *testing.T) {
	sg := storagegrouptest.Generate()

	t.Run("binary", func(t *testing.T) {
		data, err := sg.Marshal()
		require.NoError(t, err)

		sg2 := storagegroup.New()
		require.NoError(t, sg2.Unmarshal(data))

		require.Equal(t, sg, sg2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := sg.MarshalJSON()
		require.NoError(t, err)

		sg2 := storagegroup.New()
		require.NoError(t, sg2.UnmarshalJSON(data))

		require.Equal(t, sg, sg2)
	})
}

func TestNewFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *storagegroupV2.StorageGroup

		require.Nil(t, storagegroup.NewFromV2(x))
	})
}

func TestStorageGroup_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *storagegroup.StorageGroup

		require.Nil(t, x.ToV2())
	})
}

func TestNew(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		sg := storagegroup.New()

		// check initial values
		require.Nil(t, sg.Members())
		require.Nil(t, sg.ValidationDataHash())
		require.Zero(t, sg.ExpirationEpoch())
		require.Zero(t, sg.ValidationDataSize())

		// convert to v2 message
		sgV2 := sg.ToV2()

		require.Nil(t, sgV2.GetMembers())
		require.Nil(t, sgV2.GetValidationHash())
		require.Zero(t, sgV2.GetExpirationEpoch())
		require.Zero(t, sgV2.GetValidationDataSize())
	})
}
