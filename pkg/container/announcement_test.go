package container_test

import (
	"crypto/sha256"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func TestAnnouncement(t *testing.T) {
	const epoch, usedSpace uint64 = 10, 100

	cidValue := [32]byte{1, 2, 3}
	cid := container.NewID()
	cid.SetSHA256(cidValue)

	a := container.NewAnnouncement()
	a.SetEpoch(epoch)
	a.SetContainerID(cid)
	a.SetUsedSpace(usedSpace)

	require.Equal(t, epoch, a.Epoch())
	require.Equal(t, usedSpace, a.UsedSpace())
	require.Equal(t, cid, a.ContainerID())

	t.Run("test v2", func(t *testing.T) {
		const newEpoch, newUsedSpace uint64 = 20, 200

		newCidValue := [32]byte{4, 5, 6}
		newCID := new(refs.ContainerID)
		newCID.SetValue(newCidValue[:])

		v2 := a.ToV2()
		require.Equal(t, usedSpace, v2.GetUsedSpace())
		require.Equal(t, epoch, v2.GetEpoch())
		require.Equal(t, cidValue[:], v2.GetContainerID().GetValue())

		v2.SetEpoch(newEpoch)
		v2.SetUsedSpace(newUsedSpace)
		v2.SetContainerID(newCID)

		newA := container.NewAnnouncementFromV2(v2)

		require.Equal(t, newEpoch, newA.Epoch())
		require.Equal(t, newUsedSpace, newA.UsedSpace())
		require.Equal(t, container.NewIDFromV2(newCID), newA.ContainerID())
	})
}

func TestUsedSpaceEncoding(t *testing.T) {
	a := container.NewAnnouncement()
	a.SetUsedSpace(13)
	a.SetEpoch(666)

	id := container.NewID()
	id.SetSHA256([sha256.Size]byte{1, 2, 3})

	a.SetContainerID(id)

	t.Run("binary", func(t *testing.T) {
		data, err := a.Marshal()
		require.NoError(t, err)

		a2 := container.NewAnnouncement()
		require.NoError(t, a2.Unmarshal(data))

		require.Equal(t, a, a2)
	})
}
