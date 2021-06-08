package objecttest

import (
	"crypto/sha256"
	"math/rand"

	"github.com/google/uuid"
	"github.com/nspcc-dev/neofs-api-go/pkg"
	cidtest "github.com/nspcc-dev/neofs-api-go/pkg/container/id/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	ownertest "github.com/nspcc-dev/neofs-api-go/pkg/owner/test"
	sessiontest "github.com/nspcc-dev/neofs-api-go/pkg/session/test"
	refstest "github.com/nspcc-dev/neofs-api-go/pkg/test"
)

// ID returns random object.ID.
func ID() *object.ID {
	checksum := [sha256.Size]byte{}

	rand.Read(checksum[:])

	return IDWithChecksum(checksum)
}

// IDWithChecksum returns object.ID initialized
// with specified checksum.
func IDWithChecksum(cs [sha256.Size]byte) *object.ID {
	id := object.NewID()
	id.SetSHA256(cs)

	return id
}

// Address returns random object.Address.
func Address() *object.Address {
	x := object.NewAddress()

	x.SetContainerID(cidtest.Generate())
	x.SetObjectID(ID())

	return x
}

// Range returns random object.Range.
func Range() *object.Range {
	x := object.NewRange()

	x.SetOffset(1024)
	x.SetLength(2048)

	return x
}

// Attribute returns random object.Attribute.
func Attribute() *object.Attribute {
	x := object.NewAttribute()

	x.SetKey("key")
	x.SetValue("value")

	return x
}

// SplitID returns random object.SplitID.
func SplitID() *object.SplitID {
	x := object.NewSplitID()

	x.SetUUID(uuid.New())

	return x
}

func generateRaw(withParent bool) *object.RawObject {
	x := object.NewRaw()

	x.SetID(ID())
	x.SetSessionToken(sessiontest.Generate())
	x.SetPayload([]byte{1, 2, 3})
	x.SetOwnerID(ownertest.Generate())
	x.SetContainerID(cidtest.Generate())
	x.SetType(object.TypeTombstone)
	x.SetVersion(pkg.SDKVersion())
	x.SetPayloadSize(111)
	x.SetCreationEpoch(222)
	x.SetPreviousID(ID())
	x.SetParentID(ID())
	x.SetChildren(ID(), ID())
	x.SetAttributes(Attribute(), Attribute())
	x.SetSplitID(SplitID())
	x.SetPayloadChecksum(refstest.Checksum())
	x.SetPayloadHomomorphicHash(refstest.Checksum())
	x.SetSignature(refstest.Signature())

	if withParent {
		x.SetParent(generateRaw(false).Object())
	}

	return x
}

// Raw returns random object.RawObject.
func Raw() *object.RawObject {
	return generateRaw(true)
}

// Object returns random object.Object.
func Object() *object.Object {
	return Raw().Object()
}

// Tombstone returns random object.Tombstone.
func Tombstone() *object.Tombstone {
	x := object.NewTombstone()

	x.SetSplitID(SplitID())
	x.SetExpirationEpoch(13)
	x.SetMembers([]*object.ID{ID(), ID()})

	return x
}

// SplitInfo returns random object.SplitInfo.
func SplitInfo() *object.SplitInfo {
	x := object.NewSplitInfo()

	x.SetSplitID(SplitID())
	x.SetLink(ID())
	x.SetLastPart(ID())

	return x
}

// SearchFilters returns random object.SearchFilters.
func SearchFilters() object.SearchFilters {
	x := object.NewSearchFilters()

	x.AddObjectIDFilter(object.MatchStringEqual, ID())
	x.AddObjectContainerIDFilter(object.MatchStringNotEqual, cidtest.Generate())

	return x
}
