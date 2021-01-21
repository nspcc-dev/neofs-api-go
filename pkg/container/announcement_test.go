package container_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func TestAnnouncement(t *testing.T) {
	const usedSpace uint64 = 100

	cidValue := [32]byte{1, 2, 3}
	cid := container.NewID()
	cid.SetSHA256(cidValue)

	a := container.NewAnnouncement()
	a.SetContainerID(cid)
	a.SetUsedSpace(usedSpace)

	require.Equal(t, usedSpace, a.UsedSpace())
	require.Equal(t, cid, a.ContainerID())

	t.Run("test v2", func(t *testing.T) {
		const newUsedSpace uint64 = 200

		newCidValue := [32]byte{4, 5, 6}
		newCID := new(refs.ContainerID)
		newCID.SetValue(newCidValue[:])

		v2 := a.ToV2()
		require.Equal(t, usedSpace, v2.GetUsedSpace())
		require.Equal(t, cidValue[:], v2.GetContainerID().GetValue())

		v2.SetUsedSpace(newUsedSpace)
		v2.SetContainerID(newCID)

		newA := container.NewAnnouncementFromV2(v2)

		require.Equal(t, newUsedSpace, newA.UsedSpace())
		require.Equal(t, container.NewIDFromV2(newCID), newA.ContainerID())
	})
}
