package storagegroup_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/storagegroup"
	"github.com/stretchr/testify/require"
)

func TestStorageGroup_StableMarshal(t *testing.T) {
	storageGroupFrom := generateSG()

	t.Run("non empty", func(t *testing.T) {
		wire, err := storageGroupFrom.StableMarshal(nil)
		require.NoError(t, err)

		storageGroupTo := new(storagegroup.StorageGroup)
		require.NoError(t, storageGroupTo.Unmarshal(wire))

		require.Equal(t, storageGroupFrom, storageGroupTo)
	})
}

func generateChecksum(data string) *refs.Checksum {
	checksum := new(refs.Checksum)
	checksum.SetType(refs.TillichZemor)
	checksum.SetSum([]byte(data))

	return checksum
}

func generateSG() *storagegroup.StorageGroup {
	sg := new(storagegroup.StorageGroup)

	oid1 := new(refs.ObjectID)
	oid1.SetValue([]byte("Object ID 1"))

	oid2 := new(refs.ObjectID)
	oid2.SetValue([]byte("Object ID 2"))

	sg.SetValidationDataSize(300)
	sg.SetValidationHash(generateChecksum("Homomorphic hash"))
	sg.SetExpirationEpoch(100)
	sg.SetMembers([]*refs.ObjectID{oid1, oid2})

	return sg
}
