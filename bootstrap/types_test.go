package bootstrap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequestGettersSetters(t *testing.T) {
	t.Run("type", func(t *testing.T) {
		rt := NodeType(1)
		m := new(Request)

		m.SetType(rt)

		require.Equal(t, rt, m.GetType())
	})

	t.Run("state", func(t *testing.T) {
		st := Request_State(1)
		m := new(Request)

		m.SetState(st)

		require.Equal(t, st, m.GetState())
	})

	t.Run("info", func(t *testing.T) {
		info := NodeInfo{
			Address: "some address",
		}

		m := new(Request)

		m.SetInfo(info)

		require.Equal(t, info, m.GetInfo())
	})
}
