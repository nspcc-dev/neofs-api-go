package containertest

import (
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	cidtest "github.com/nspcc-dev/neofs-api-go/pkg/container/id/test"
	netmaptest "github.com/nspcc-dev/neofs-api-go/pkg/netmap/test"
	ownertest "github.com/nspcc-dev/neofs-api-go/pkg/owner/test"
	refstest "github.com/nspcc-dev/neofs-api-go/pkg/test"
)

// Attribute returns random container.Attribute.
func Attribute() *container.Attribute {
	x := container.NewAttribute()

	x.SetKey("key")
	x.SetValue("value")

	return x
}

// Attributes returns random container.Attributes.
func Attributes() container.Attributes {
	return container.Attributes{Attribute(), Attribute()}
}

// Container returns random container.Container.
func Container() *container.Container {
	x := container.New()

	x.SetVersion(refstest.Version())
	x.SetAttributes(Attributes())
	x.SetOwnerID(ownertest.Generate())
	x.SetBasicACL(123)
	x.SetPlacementPolicy(netmaptest.PlacementPolicy())

	return x
}

// UsedSpaceAnnouncement returns random container.UsedSpaceAnnouncement.
func UsedSpaceAnnouncement() *container.UsedSpaceAnnouncement {
	x := container.NewAnnouncement()

	x.SetContainerID(cidtest.Generate())
	x.SetEpoch(55)
	x.SetUsedSpace(999)

	return x
}
