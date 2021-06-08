package storagegrouptest

import (
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	objecttest "github.com/nspcc-dev/neofs-api-go/pkg/object/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/storagegroup"
	refstest "github.com/nspcc-dev/neofs-api-go/pkg/test"
)

// Generate returns random storagegroup.StorageGroup.
func Generate() *storagegroup.StorageGroup {
	x := storagegroup.New()

	x.SetExpirationEpoch(66)
	x.SetValidationDataSize(322)
	x.SetValidationDataHash(refstest.Checksum())
	x.SetMembers([]*object.ID{objecttest.ID(), objecttest.ID()})

	return x
}
