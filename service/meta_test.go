package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCutRestoreMeta(t *testing.T) {
	items := []func() SeizedMetaHeaderContainer{
		func() SeizedMetaHeaderContainer {
			m := new(RequestMetaHeader)
			m.SetEpoch(1)
			return m
		},
	}

	for _, item := range items {
		v1 := item()
		m1 := v1.CutMeta()
		v1.RestoreMeta(m1)

		require.Equal(t, item(), v1)
	}
}
