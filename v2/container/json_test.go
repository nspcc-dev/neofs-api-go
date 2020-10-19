package container_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/container"
	"github.com/stretchr/testify/require"
)

func TestContainerJSON(t *testing.T) {
	exp := generateContainer("container")

	t.Run("non empty", func(t *testing.T) {
		data := container.ContainerToJSON(exp)
		require.NotNil(t, data)

		got := container.ContainerFromJSON(data)
		require.NotNil(t, got)

		require.Equal(t, exp, got)
	})
}
