package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSetRaw(t *testing.T) {
	items := []RawContainer{
		new(RequestMetaHeader),
	}

	for _, item := range items {
		// init with false
		item.SetRaw(false)

		item.SetRaw(true)
		require.True(t, item.GetRaw())

		item.SetRaw(false)
		require.False(t, item.GetRaw())
	}
}
