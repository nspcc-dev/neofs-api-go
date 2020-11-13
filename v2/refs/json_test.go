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
