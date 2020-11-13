package storagegroup_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/storagegroup"
	"github.com/stretchr/testify/require"
)

func TestStorageGroupJSON(t *testing.T) {
	sg := generateSG()

	data, err := sg.MarshalJSON()
	require.NoError(t, err)

	sg2 := new(storagegroup.StorageGroup)
	require.NoError(t, sg2.UnmarshalJSON(data))

	require.Equal(t, sg, sg2)
}
