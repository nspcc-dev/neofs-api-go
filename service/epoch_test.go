package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSetEpoch(t *testing.T) {
	v := uint64(5)

	items := []EpochContainer{
		new(ResponseMetaHeader),
		new(RequestMetaHeader),
	}

	for _, item := range items {
		item.SetEpoch(v)
		require.Equal(t, v, item.GetEpoch())
	}
}
