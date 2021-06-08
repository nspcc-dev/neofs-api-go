package object

import (
	"strings"
	"testing"

	cidtest "github.com/nspcc-dev/neofs-api-go/pkg/container/id/test"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func TestAddress_SetContainerID(t *testing.T) {
	a := NewAddress()

	id := cidtest.Generate()

	a.SetContainerID(id)

	require.Equal(t, id, a.ContainerID())
}

func TestAddress_SetObjectID(t *testing.T) {
	a := NewAddress()

	oid := randID(t)

	a.SetObjectID(oid)

	require.Equal(t, oid, a.ObjectID())
}

func TestAddress_Parse(t *testing.T) {
	cid := cidtest.Generate()

	oid := NewID()
	oid.SetSHA256(randSHA256Checksum(t))

	t.Run("should parse successful", func(t *testing.T) {
		s := strings.Join([]string{cid.String(), oid.String()}, addressSeparator)
		a := NewAddress()

		require.NoError(t, a.Parse(s))
		require.Equal(t, oid, a.ObjectID())
		require.Equal(t, cid, a.ContainerID())
	})

	t.Run("should fail for bad address", func(t *testing.T) {
		s := strings.Join([]string{cid.String()}, addressSeparator)
		require.EqualError(t, NewAddress().Parse(s), errInvalidAddressString.Error())
	})

	t.Run("should fail on container.ID", func(t *testing.T) {
		s := strings.Join([]string{"1", "2"}, addressSeparator)
		require.Error(t, NewAddress().Parse(s))
	})

	t.Run("should fail on object.ID", func(t *testing.T) {
		s := strings.Join([]string{cid.String(), "2"}, addressSeparator)
		require.Error(t, NewAddress().Parse(s))
	})
}

func TestAddressEncoding(t *testing.T) {
	a := NewAddress()
	a.SetObjectID(randID(t))
	a.SetContainerID(cidtest.Generate())

	t.Run("binary", func(t *testing.T) {
		data, err := a.Marshal()
		require.NoError(t, err)

		a2 := NewAddress()
		require.NoError(t, a2.Unmarshal(data))

		require.Equal(t, a, a2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := a.MarshalJSON()
		require.NoError(t, err)

		a2 := NewAddress()
		require.NoError(t, a2.UnmarshalJSON(data))

		require.Equal(t, a, a2)
	})
}

func TestNewAddressFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *refs.Address

		require.Nil(t, NewAddressFromV2(x))
	})
}

func TestAddress_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *Address

		require.Nil(t, x.ToV2())
	})
}
