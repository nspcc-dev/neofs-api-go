package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockedRequest struct {
	msg  string
	name string
	role NodeRole
	code codes.Code
	RequestMetaHeader
}

func TestMetaRequest(t *testing.T) {
	tests := []mockedRequest{
		{
			role:              InnerRingNode,
			name:              "direct to ir node",
			RequestMetaHeader: RequestMetaHeader{TTL: NonForwardingTTL},
		},
		{
			role:              StorageNode,
			code:              codes.InvalidArgument,
			msg:               ErrIncorrectTTL.Error(),
			name:              "direct to storage node",
			RequestMetaHeader: RequestMetaHeader{TTL: NonForwardingTTL},
		},
		{
			role:              StorageNode,
			msg:               ErrZeroTTL.Error(),
			code:              codes.InvalidArgument,
			name:              "zero ttl",
			RequestMetaHeader: RequestMetaHeader{TTL: ZeroTTL},
		},
		{
			role:              InnerRingNode,
			name:              "default to ir node",
			RequestMetaHeader: RequestMetaHeader{TTL: SingleForwardingTTL},
		},
		{
			role:              StorageNode,
			name:              "default to storage node",
			RequestMetaHeader: RequestMetaHeader{TTL: SingleForwardingTTL},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			before := tt.GetTTL()
			err := ProcessRequestTTL(&tt, IRNonForwarding(tt.role))
			if tt.msg != "" {
				require.Errorf(t, err, tt.msg)

				state, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, state.Code(), tt.code)
				require.Equal(t, state.Message(), tt.msg)
			} else {
				require.NoError(t, err)
				require.NotEqualf(t, before, tt.GetTTL(), "ttl should be changed: %d vs %d", before, tt.GetTTL())
			}
		})
	}
}
