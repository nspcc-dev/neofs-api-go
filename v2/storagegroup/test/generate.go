package storagegrouptest

import (
	refstest "github.com/nspcc-dev/neofs-api-go/v2/refs/test"
	"github.com/nspcc-dev/neofs-api-go/v2/storagegroup"
)

func GenerateStorageGroup(empty bool) *storagegroup.StorageGroup {
	m := new(storagegroup.StorageGroup)

	if !empty {
		m.SetValidationDataSize(44)
		m.SetExpirationEpoch(55)
	}

	m.SetValidationHash(refstest.GenerateChecksum(empty))
	m.SetMembers(refstest.GenerateObjectIDs(empty))

	return m
}
