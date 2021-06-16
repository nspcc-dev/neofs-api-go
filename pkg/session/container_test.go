package session_test

import (
	"testing"

	cidtest "github.com/nspcc-dev/neofs-api-go/pkg/container/id/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/session"
	sessiontest "github.com/nspcc-dev/neofs-api-go/pkg/session/test"
	v2session "github.com/nspcc-dev/neofs-api-go/v2/session"
	"github.com/stretchr/testify/require"
)

func TestContainerContextVerbs(t *testing.T) {
	c := session.NewContainerContext()

	assert := func(setter func(), getter func() bool, verb v2session.ContainerSessionVerb) {
		setter()

		require.True(t, getter())

		require.Equal(t, verb, c.ToV2().Verb())
	}

	t.Run("PUT", func(t *testing.T) {
		assert(c.ForPut, c.IsForPut, v2session.ContainerVerbPut)
	})

	t.Run("DELETE", func(t *testing.T) {
		assert(c.ForDelete, c.IsForDelete, v2session.ContainerVerbDelete)
	})

	t.Run("SETEACL", func(t *testing.T) {
		assert(c.ForSetEACL, c.IsForSetEACL, v2session.ContainerVerbSetEACL)
	})
}

func TestContainerContext_ApplyTo(t *testing.T) {
	c := session.NewContainerContext()
	id := cidtest.Generate()

	t.Run("method", func(t *testing.T) {
		c.ApplyTo(id)

		require.Equal(t, id, c.Container())

		c.ApplyTo(nil)

		require.Nil(t, c.Container())
	})

	t.Run("helper functions", func(t *testing.T) {
		c.ApplyTo(id)

		session.ApplyToAllContainers(c)

		require.Nil(t, c.Container())
	})
}

func TestFilter_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *session.ContainerContext

		require.Nil(t, x.ToV2())
	})

	t.Run("default values", func(t *testing.T) {
		c := session.NewContainerContext()

		// check initial values
		require.Nil(t, c.Container())

		for _, op := range []func() bool{
			c.IsForPut,
			c.IsForDelete,
			c.IsForSetEACL,
		} {
			require.False(t, op())
		}

		// convert to v2 message
		cV2 := c.ToV2()

		require.Equal(t, v2session.ContainerVerbUnknown, cV2.Verb())
		require.True(t, cV2.Wildcard())
		require.Nil(t, cV2.ContainerID())
	})
}

func TestContainerContextEncoding(t *testing.T) {
	c := sessiontest.ContainerContext()

	t.Run("binary", func(t *testing.T) {
		data, err := c.Marshal()
		require.NoError(t, err)

		c2 := session.NewContainerContext()
		require.NoError(t, c2.Unmarshal(data))

		require.Equal(t, c, c2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := c.MarshalJSON()
		require.NoError(t, err)

		c2 := session.NewContainerContext()
		require.NoError(t, c2.UnmarshalJSON(data))

		require.Equal(t, c, c2)
	})
}
