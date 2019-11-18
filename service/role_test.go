package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNodeRole_String(t *testing.T) {
	tests := []struct {
		nt   NodeRole
		want string
	}{
		{want: "Unknown"},
		{nt: StorageNode, want: "StorageNode"},
		{nt: InnerRingNode, want: "InnerRingNode"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			require.Equal(t, tt.want, tt.nt.String())
		})
	}
}
