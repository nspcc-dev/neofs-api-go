package object

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/stretchr/testify/require"
)

func TestAddress_SetContainerID(t *testing.T) {
	a := NewAddress()

	cid := container.NewID()
	cid.SetSHA256(randSHA256Checksum(t))

	a.SetContainerID(cid)

	require.Equal(t, cid, a.GetContainerID())
}

func TestAddress_SetObjectID(t *testing.T) {
	a := NewAddress()

	oid := randID(t)

	a.SetObjectID(oid)

	require.Equal(t, oid, a.GetObjectID())
}
