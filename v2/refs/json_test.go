package refs_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func TestAddressJSON(t *testing.T) {
	a := generateAddress([]byte{1}, []byte{2})

	data, err := a.MarshalJSON()
	require.NoError(t, err)

	a2 := new(refs.Address)
	require.NoError(t, a2.UnmarshalJSON(data))

	require.Equal(t, a, a2)
}

func TestObjectIDJSON(t *testing.T) {
	o := new(refs.ObjectID)
	o.SetValue([]byte{1})

	data, err := o.MarshalJSON()
	require.NoError(t, err)

	o2 := new(refs.ObjectID)
	require.NoError(t, o2.UnmarshalJSON(data))

	require.Equal(t, o, o2)
}

func TestContainerIDJSON(t *testing.T) {
	cid := new(refs.ContainerID)
	cid.SetValue([]byte{1})

	data, err := cid.MarshalJSON()
	require.NoError(t, err)

	cid2 := new(refs.ContainerID)
	require.NoError(t, cid2.UnmarshalJSON(data))

	require.Equal(t, cid, cid2)
}

func TestOwnerIDJSON(t *testing.T) {
	o := new(refs.OwnerID)
	o.SetValue([]byte{1})

	data, err := o.MarshalJSON()
	require.NoError(t, err)

	o2 := new(refs.OwnerID)
	require.NoError(t, o2.UnmarshalJSON(data))

	require.Equal(t, o, o2)
}

func TestVersionSON(t *testing.T) {
	v := generateVersion(1, 2)

	data, err := v.MarshalJSON()
	require.NoError(t, err)

	v2 := new(refs.Version)
	require.NoError(t, v2.UnmarshalJSON(data))

	require.Equal(t, v, v2)
}
