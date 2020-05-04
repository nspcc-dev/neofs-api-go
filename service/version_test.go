package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSetVersion(t *testing.T) {
	v := uint32(7)

	items := []VersionContainer{
		new(ResponseMetaHeader),
		new(RequestMetaHeader),
	}

	for _, item := range items {
		item.SetVersion(v)
		require.Equal(t, v, item.GetVersion())
	}
}
