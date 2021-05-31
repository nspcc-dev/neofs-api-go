package container

import (
	cid "github.com/nspcc-dev/neofs-api-go/pkg/container/id"
	"github.com/nspcc-dev/neofs-api-go/v2/container"
)

// UsedSpaceAnnouncement is an announcement message used by storage nodes to
// estimate actual container sizes.
type UsedSpaceAnnouncement container.UsedSpaceAnnouncement

// NewAnnouncement initialize empty UsedSpaceAnnouncement message.
func NewAnnouncement() *UsedSpaceAnnouncement {
	return NewAnnouncementFromV2(new(container.UsedSpaceAnnouncement))
}

// NewAnnouncementFromV2 wraps protocol dependent version of
// UsedSpaceAnnouncement message.
func NewAnnouncementFromV2(v *container.UsedSpaceAnnouncement) *UsedSpaceAnnouncement {
	return (*UsedSpaceAnnouncement)(v)
}

// Epoch of the announcement.
func (a *UsedSpaceAnnouncement) Epoch() uint64 {
	return (*container.UsedSpaceAnnouncement)(a).GetEpoch()
}

// SetEpoch sets announcement epoch value.
func (a *UsedSpaceAnnouncement) SetEpoch(epoch uint64) {
	(*container.UsedSpaceAnnouncement)(a).SetEpoch(epoch)
}

// ContainerID of the announcement.
func (a *UsedSpaceAnnouncement) ContainerID() *cid.ID {
	return cid.NewFromV2(
		(*container.UsedSpaceAnnouncement)(a).GetContainerID(),
	)
}

// SetContainerID sets announcement container value.
func (a *UsedSpaceAnnouncement) SetContainerID(cid *cid.ID) {
	(*container.UsedSpaceAnnouncement)(a).SetContainerID(cid.ToV2())
}

// UsedSpace in container.
func (a *UsedSpaceAnnouncement) UsedSpace() uint64 {
	return (*container.UsedSpaceAnnouncement)(a).GetUsedSpace()
}

// SetUsedSpace sets used space value by specified container.
func (a *UsedSpaceAnnouncement) SetUsedSpace(value uint64) {
	(*container.UsedSpaceAnnouncement)(a).SetUsedSpace(value)
}

// ToV2 returns protocol dependent version of UsedSpaceAnnouncement message.
func (a *UsedSpaceAnnouncement) ToV2() *container.UsedSpaceAnnouncement {
	return (*container.UsedSpaceAnnouncement)(a)
}

// Marshal marshals UsedSpaceAnnouncement into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (a *UsedSpaceAnnouncement) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return a.ToV2().
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of UsedSpaceAnnouncement.
func (a *UsedSpaceAnnouncement) Unmarshal(data []byte) error {
	return a.ToV2().
		Unmarshal(data)
}
