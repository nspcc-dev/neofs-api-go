package object

import (
	"strings"
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

func TestAddress_Parse(t *testing.T) {
	cid := container.NewID()
	cid.SetSHA256(randSHA256Checksum(t))

	oid := NewID()
	oid.SetSHA256(randSHA256Checksum(t))

	t.Run("should parse successful", func(t *testing.T) {
		s := strings.Join([]string{cid.String(), oid.String()}, addressSeparator)
		a := NewAddress()

		require.NoError(t, a.Parse(s))
		require.Equal(t, oid, a.GetObjectID())
		require.Equal(t, cid, a.GetContainerID())
	})

	t.Run("should fail for bad address", func(t *testing.T) {
		s := strings.Join([]string{cid.String()}, addressSeparator)
		require.EqualError(t, NewAddress().Parse(s), ErrBadAddress.Error())
	})

	t.Run("should fail on container.ID", func(t *testing.T) {
		s := strings.Join([]string{"1", "2"}, addressSeparator)
		require.EqualError(t, NewAddress().Parse(s), container.ErrBadID.Error())
	})

	t.Run("should fail on object.ID", func(t *testing.T) {
		s := strings.Join([]string{cid.String(), "2"}, addressSeparator)
		require.EqualError(t, NewAddress().Parse(s), ErrBadID.Error())
	})
}
