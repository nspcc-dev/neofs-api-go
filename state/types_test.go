package state

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChangeStateRequestGettersSetters(t *testing.T) {
	t.Run("state", func(t *testing.T) {
		st := ChangeStateRequest_State(1)
		m := new(ChangeStateRequest)

		m.SetState(st)

		require.Equal(t, st, m.GetState())
	})
}
