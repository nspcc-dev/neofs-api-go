package container_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/container"
	"github.com/stretchr/testify/require"
)

func TestContainerJSON(t *testing.T) {
	exp := generateContainer("container")

	t.Run("non empty", func(t *testing.T) {
		data, err := container.ContainerToJSON(exp)
		require.NoError(t, err)

		got, err := container.ContainerFromJSON(data)
		require.NoError(t, err)

		require.Equal(t, exp, got)
	})

	t.Run("empty", func(t *testing.T) {
		_, err := container.ContainerToJSON(nil)
		require.Error(t, err)

		_, err = container.ContainerFromJSON(nil)
		require.Error(t, err)
	})
}
