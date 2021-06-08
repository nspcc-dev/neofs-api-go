package storagegroup_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	objecttest "github.com/nspcc-dev/neofs-api-go/pkg/object/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/storagegroup"
	storagegrouptest "github.com/nspcc-dev/neofs-api-go/pkg/storagegroup/test"
	refstest "github.com/nspcc-dev/neofs-api-go/pkg/test"
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
