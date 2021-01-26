package container

import (
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
func (a *UsedSpaceAnnouncement) ContainerID() *ID {
	return NewIDFromV2(
		(*container.UsedSpaceAnnouncement)(a).GetContainerID(),
	)
}

// SetContainerID sets announcement container value.
func (a *UsedSpaceAnnouncement) SetContainerID(cid *ID) {
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
